/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-27 10:42:42
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-27 10:42:49
 * @FilePath: /xoj-backend/model/entity/User.go
 * @Description: 实体类，表字段
 */
package entity

import (
	"time"
)

// 用户
type User struct {
	// id
	ID int64 `json:"id"`
	// 账号
	Useraccount string `json:"useraccount"`
	// 密码
	Userpassword string `json:"userpassword"`
	// 微信开放平台id
	Unionid string `json:"unionid"`
	// 公众号openId
	Mpopenid string `json:"mpopenid"`
	// 用户昵称
	Username string `json:"username"`
	// 用户头像
	Useravatar string `json:"useravatar"`
	// 用户简介
	Userprofile string `json:"userprofile"`
	// 性别
	Gender int32 `json:"gender"`
	// 用户角色：user/admin/ban
	Userrole string `json:"userrole"`
	// 创建时间
	Createtime time.Time `json:"createtime"`
	// 更新时间
	Updatetime time.Time `json:"updatetime"`
	// 是否删除
	Isdelete int32 `json:"isdelete"`
}
