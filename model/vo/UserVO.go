/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-27 10:39:07
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-27 10:41:25
 * @FilePath: /xoj-backend/model/vo/UserVO.go
 * @Description: 专门返回给前端用的，可以节约网络传输大小，或者过滤字段（脱敏）、保证安全性。
 */
package vo

import (
	"time"
)

type UserVO struct {
	// id
	ID int64 `json:"id"`
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
}
