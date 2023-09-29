/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-27 10:44:32
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-29 11:44:29
 * @FilePath: /xoj-backend/model/dto/user/UserAddRequest.go
 * @Description: 业务层面的封装
 */
package user

// 用户登录请求
type UserLoginRequest struct {
	// 账号
	UserAccount string `json:"userAccount" form:"userAccount"`
	// 密码
	UserPassword string `json:"userPassword" form:"userPassword"`
}
