/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-26 10:35:03
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-27 11:22:02
 * @FilePath: /xoj-backend/models/user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
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
