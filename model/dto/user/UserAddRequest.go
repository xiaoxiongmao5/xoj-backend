/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-27 10:44:32
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-29 11:44:25
 * @FilePath: /xoj-backend/model/dto/user/UserAddRequest.go
 * @Description: 业务层面的封装
 */
package user

// 用户创建请求
type UserAddRequest struct {
	// 账号
	UserAccount string `json:"userAccount"`
	// 密码
	UserPassword string `json:"userPassword"`
	// 用户昵称
	UserName string `json:"userName"`
	// 用户角色：user/admin/ban
	UserRole string `json:"userRole"`
}
