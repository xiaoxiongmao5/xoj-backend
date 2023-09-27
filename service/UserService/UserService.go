package userservice

import (
	"context"
	"database/sql"
	"errors"

	beecontext "github.com/beego/beego/v2/server/web/context"
	"github.com/dgrijalva/jwt-go"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/dbsq"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/dto/user"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/enums"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mydb"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mylog"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/myresq"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/store"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/utils"
)

// 获取当前登录用户
func GetLoginUser(ctx *beecontext.Context) *dbsq.User {
	// 先判断是否已登录
	// 从请求中获取当前的 Token
	tokenCookie := ctx.GetCookie("token") //todo USER_LOGIN_STATE
	if tokenCookie == "" {
		myresq.Abort(ctx, myresq.NOT_LOGIN_ERROR, "")
		return nil
	}

	// 验证当前 Token
	token, err := jwt.Parse(tokenCookie, func(token *jwt.Token) (interface{}, error) {
		return []byte("SecretKey"), nil
	})
	if err != nil || !token.Valid {
		myresq.Abort(ctx, myresq.NOT_LOGIN_ERROR, "")
		return nil
	}
	// 从 Token 中获取用户信息
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		myresq.Abort(ctx, myresq.NOT_LOGIN_ERROR, "")
		return nil
	}

	// 重新生成 Token，并更新有效期
	userAccount := claims["user_account"].(string)
	userRole := claims["user_role"].(string)
	if newToken, err := utils.GenerateToken(userAccount, userRole); err != nil {
		mylog.Log.Warn("重新生成 Token 以更新有效期失败 err=", err.Error())
	} else {
		// 删除旧的 token
		delete(store.TokenMemoryStore, tokenCookie)

		// 更新内存中的 token 数据
		store.TokenMemoryStore[newToken] = true

		// 将新的 token 返回给前端
		domain, _ := utils.GetDomainFromReferer(ctx.Request.Referer())
		ctx.SetCookie("token", newToken, 3600, "/", domain, false, true)
	}

	// 从数据库查询（追求性能的话可以注释，直接走缓存）
	loginUser, err := GetUserInfoByUserAccount(userAccount)
	if err != nil {
		mylog.Log.Errorf("q.GetUserInfoByUserAccount err=%v", err.Error())
		myresq.Abort(ctx, myresq.USER_NOT_EXIST, "")
		return nil
	}
	return loginUser
}

// 是否为管理员
func IsAdmin(user *dbsq.User) bool {
	return utils.CheckSame[string]("是否为管理员", user.Userrole, enums.ADMIN.GetValue())
}

// 根据Id 获取用户信息
func GetUserInfoById(id int64) (*dbsq.User, error) {
	conn, err := mydb.GetConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	q := dbsq.New(conn)
	ctx := context.Background()
	return q.GetUserInfoById(ctx, id)
}

// 根据userAccount 获得用户信息
func GetUserInfoByUserAccount(userAccount string) (*dbsq.User, error) {
	conn, err := mydb.GetConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	q := dbsq.New(conn)
	ctx := context.Background()
	return q.GetUserInfoByUniUserAccount(ctx, userAccount)
}

// 创建账号
func CreateUser(param *user.UserRegisterRequest) (sql.Result, error) {
	userAccount, userPassword, checkUserPassword := param.Useraccount, param.Userpassword, param.CheckUserpassword
	// 检验
	if utils.IsAnyBlank(userAccount, userPassword, checkUserPassword) {
		return nil, errors.New("参数为空")
	}
	if length := len(userAccount); length < 4 || length > 16 {
		return nil, errors.New("用户账号长度不符合规定,长度要求4~16位")
	}
	if length := len(userPassword); length < 6 || length > 16 {
		return nil, errors.New("用户密码长度不符合规定,长度要求6~16位")
	}
	// 密码和校验密码相同
	if !utils.CheckSame[string]("校验两次输入的密码一致", userPassword, checkUserPassword) {
		return nil, errors.New("两次输入的密码不一致")
	}
	// 账号不能重复
	if _, err := GetUserInfoByUserAccount(userAccount); err == nil {
		return nil, errors.New("账户已存在")
	}
	// 将密码进行哈希
	hashPassword, err := utils.HashPasswordByBcrypt(userPassword)
	if err != nil {
		mylog.Log.Errorf("utils.HashByBcrypt err=%v", err.Error())
		return nil, err
	}

	// 将 userAccount 转换为 sql.NullString
	var userName sql.NullString
	if userAccount != "" {
		userName = sql.NullString{String: userAccount, Valid: true}
	} else {
		userName = sql.NullString{String: "", Valid: false}
	}

	params := &dbsq.CreateUserParams{
		Useraccount:  userAccount,
		Userpassword: hashPassword,
		Username:     userName,
	}
	conn, err := mydb.GetConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	q := dbsq.New(conn)
	ctx := context.Background()
	return q.CreateUser(ctx, params)
}

// // 删除token
// func DeleteToken(c *gin.Context) {
// 	// 从cookie拿到token
// 	tokenCookie, err := c.Cookie("token")
// 	if err != nil || tokenCookie == "" {
// 		return
// 	}

// 	// 从服务端删除该token
// 	delete(mylog.TokenMemoryStore, tokenCookie)
// }
