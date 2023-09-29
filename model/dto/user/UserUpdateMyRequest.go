/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-27 10:44:32
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-27 11:06:19
 * @FilePath: /xoj-backend/model/dto/user/UserAddRequest.go
 * @Description: 业务层面的封装
 */
package user

// 用户更新个人信息请求
type UserUpdateMyRequest struct {
	// 用户昵称
	UserName string `json:"userName"`
	// 用户头像
	UserAvatar string `json:"userAvatar"`
	// 用户简介
	UserProfile string `json:"userProfile"`
}
