package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/xiaoxiongmao5/xoj/xoj-backend/dbsq"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/models"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mydb"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mylog"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/utils"
)

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
func CreateUser(param *models.UserRegisterParams) (sql.Result, error) {
	userAccount, userPassword, checkUserPassword := param.UserAccount, param.UserPassword, param.CheckUserPassword
	// 检验
	if utils.AreEmptyStrings(userAccount, userPassword, checkUserPassword) {
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
