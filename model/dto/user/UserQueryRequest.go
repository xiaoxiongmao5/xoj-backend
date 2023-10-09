/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-27 10:44:32
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-28 18:50:13
 * @FilePath: /xoj-backend/model/dto/user/UserAddRequest.go
 * @Description: 业务层面的封装
 */
package user

import "github.com/xiaoxiongmao5/xoj/xoj-backend/model/common"

// 用户查询请求
type UserQueryRequest struct {
	common.PageRequest
	// id
	Id int64 `json:"id"`
	// 微信开放平台id
	UnionId string `json:"unionId"`
	// 公众号openId
	MpOpenId string `json:"mpOpenId"`
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
}
