/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-27 10:39:07
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-29 14:15:17
 * @FilePath: /xoj-backend/model/vo/UserVO.go
 * @Description: 专门返回给前端用的，可以节约网络传输大小，或者过滤字段（脱敏）、保证安全性。
 */
package vo

import (
	"time"
)

type LoginUserVO struct {
	// id
	ID int64 `json:"id"`
	// 用户昵称
	UserName string `json:"userName"`
	// 用户头像
	UserAvatar string `json:"userAvatar"`
	// 用户简介
	UserProfile string `json:"userProfile"`
	// 性别
	Gender int32 `json:"gender"`
	// 用户角色：user/admin/ban
	UserRole string `json:"userRole"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
}
