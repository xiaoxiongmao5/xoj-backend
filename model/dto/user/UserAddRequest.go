/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-27 10:44:32
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-27 11:01:05
 * @FilePath: /xoj-backend/model/dto/user/UserAddRequest.go
 * @Description: 业务层面的封装
 */
package user

// 用户创建请求
type UserAddRequest struct {
	// 账号
	Useraccount string `json:"useraccount"`
	// 密码
	Userpassword string `json:"userpassword"`
	// 用户昵称
	Username string `json:"username"`
	// 用户角色：user/admin/ban
	Userrole string `json:"userrole"`
}
