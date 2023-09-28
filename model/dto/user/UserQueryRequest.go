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
	ID int64 `json:"id"`
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
}
