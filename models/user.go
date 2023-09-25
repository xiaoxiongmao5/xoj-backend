package models

import (
	"time"

	"github.com/xiaoxiongmao5/xoj/xoj-backend/dbsq"
)

// 登录获取用户信息显示
type ShowUserJSON struct {
	// id
	ID int64 `json:"id"`
	// 用户昵称
	Username string `json:"username"`
	// 账号
	Useraccount string `json:"useraccount"`
	// 用户头像
	Useravatar string `json:"useravatar"`
	// 性别
	Gender int32 `json:"gender"`
	// 用户角色：user / admin
	Userrole string `json:"userrole"`
	// 创建时间
	Createtime time.Time `json:"createtime"`
	// 更新时间
	Updatetime time.Time `json:"updatetime"`
}

// 注册用户
type UserRegisterParams struct {
	UserAccount       string `json:"useraccount"`
	UserPassword      string `json:"userpassword"`
	CheckUserPassword string `json:"checkUserpassword"`
}

type UserLoginParams struct {
	Useraccount  string `json:"useraccount" form:"useraccount"`
	Userpassword string `json:"userpassword" form:"userpassword"`
}

func ConvertToNormalUser(u *dbsq.User) *ShowUserJSON {
	nu := &ShowUserJSON{
		ID:          u.ID,
		Useraccount: u.Useraccount,
		Userrole:    u.Userrole,
		Createtime:  u.Createtime,
		Updatetime:  u.Updatetime,
	}

	if u.Username.Valid {
		nu.Username = u.Username.String
	}

	if u.Useravatar.Valid {
		nu.Useravatar = u.Useravatar.String
	}

	if u.Gender.Valid {
		nu.Gender = u.Gender.Int32
	}

	return nu
}
