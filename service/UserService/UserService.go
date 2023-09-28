package userservice

import (
	"context"
	"database/sql"
	"errors"

	"github.com/beego/beego/v2/client/orm"
	beecontext "github.com/beego/beego/v2/server/web/context"
	"github.com/dgrijalva/jwt-go"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/constant"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/dbsq"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/dto/user"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/entity"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/enums"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/vo"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mydb"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mylog"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/myresq"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/store"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/utils"
)

func UserRegister(useraccount, userpassword, checkUserPassword string) {

}

func UserLogin(useraccount, userpassword string) {

}

func UserLoginByMpOpen(wxOAuth2UserInfo interface{}) {

}

// 获取当前登录用户
//
//	@param ctx
//	@return *dbsq.User
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

// 获取当前登录用户（允许未登录）
//
//	@param ctx
func GetLoginUserPermitNull(ctx *beecontext.Context) {

}

// 是否为管理员
//
//	@param user
//	@return bool
func IsAdmin(user *dbsq.User) bool {
	return utils.CheckSame[string]("是否为管理员", user.Userrole, enums.ADMIN.GetValue())
}

// 用户注销
func UserLogout() {

}

// func GetLoginUserVO(userinfo entity.User) vo.LoginUserVO {

// }

func GetUserVO(original *entity.User) vo.UserVO {
	var userVO vo.UserVO
	utils.CopyStructFields(original, &userVO)
	return userVO
}

func GetUserVO2(original *dbsq.User) vo.UserVO {
	var userVO vo.UserVO
	utils.CopyStructFields(original, &userVO)
	return userVO
}

// ListUserVO
//
//	@param list
//	@return respdata
func ListUserVO(list []*entity.User) (respdata []vo.UserVO) {
	if utils.IsEmpty(list) {
		return
	}
	for _, one := range list {
		respdata = append(respdata, GetUserVO(one))
	}
	return
}

// GetQuerySeter
//
//	@param qs
//	@param queryRequest
//	@return orm.QuerySeter
func GetQuerySeter(qs orm.QuerySeter, queryRequest user.UserQueryRequest) orm.QuerySeter {
	id := queryRequest.ID
	unionid := queryRequest.Unionid
	mpopenid := queryRequest.Mpopenid
	username := queryRequest.Username
	userprofile := queryRequest.Userprofile
	userrole := queryRequest.Userrole
	sortField := queryRequest.PageRequest.SortField
	sortOrder := queryRequest.PageRequest.SortOrder

	if id != 0 {
		qs = qs.Filter("id", id)
	}
	if utils.IsNotBlank(unionid) {
		qs = qs.Filter("unionId", unionid)
	}
	if utils.IsNotBlank(mpopenid) {
		qs = qs.Filter("mpOpenId", mpopenid)
	}
	if utils.IsNotBlank(userrole) {
		qs = qs.Filter("userRole", userrole)
	}
	if utils.IsNotBlank(userprofile) {
		qs = qs.Filter("userProfile__icontains", userprofile)
	}
	if utils.IsNotBlank(username) {
		qs = qs.Filter("userName__icontains", username)
	}

	if utils.IsNotBlank(sortField) {
		order := sortField
		if utils.CheckSame[string]("检查排序是否一样", sortOrder, constant.SORT_ORDER_DESC) {
			order = "-" + order
		}
		qs = qs.OrderBy(order)
	}
	return qs
}

// ListByIds
//
//	@param ids
//	@return []*entity.User
//	@return error
func ListByIds(ids []int64) ([]*entity.User, error) {
	qs := mydb.O.QueryTable(new(entity.User))
	qs = qs.Filter("id__in", ids).Filter("isDelete", 0)
	var users []*entity.User
	_, err := qs.All(&users)
	if err != nil {
		mylog.Log.Errorf("User ListByIds qs.All error: %v", err.Error())
		return users, err
	}
	return users, nil
}

// GetById
//
//	@param id
//	@return *entity.User
//	@return error
func GetById(id int64) (*entity.User, error) {
	qs := mydb.O.QueryTable(new(entity.User))
	var userInfo entity.User
	err := qs.Filter("id", id).Filter("isDelete", 0).One(&userInfo)
	if err == orm.ErrMultiRows {
		mylog.Log.Errorf("user 表中存在 id=[%d] 的多条记录, qs.One err=[%v]", id, err.Error())
		return nil, err
	}
	if err == orm.ErrNoRows {
		mylog.Log.Errorf("user 表没有找到 id=[%d] 的记录, qs.One err=[%v]", id, err.Error())
		return nil, err
	}
	return &userInfo, nil
}

func GetById2(id int64) (*dbsq.User, error) {
	conn, err := mydb.GetConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	q := dbsq.New(conn)
	ctx := context.Background()
	return q.GetUserInfoById(ctx, id)
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
