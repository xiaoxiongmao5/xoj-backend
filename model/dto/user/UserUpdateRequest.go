/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-27 10:44:32
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-27 11:06:57
 * @FilePath: /xoj-backend/model/dto/user/UserAddRequest.go
 * @Description: 业务层面的封装
 */
package user

// 用户更新请求
type UserUpdateRequest struct {
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
}
