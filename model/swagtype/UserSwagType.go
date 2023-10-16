/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-30 11:37:36
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-16 20:33:58
 */
package swagtype

import (
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/entity"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/vo"
)

type UserResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    entity.User `json:"data"`
}

type UserVOResponse struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    vo.UserVO `json:"data"`
}

type LoginUserVOResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    vo.LoginUserVO `json:"data"`
}

type ListUserResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    listUser `json:"data"`
}
type listUser struct {
	Records []entity.User `json:"records"`
	Total   int64         `json:"total"`
}

type ListUserVOResponse struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    listUserVO `json:"data"`
}
type listUserVO struct {
	Records []vo.UserVO `json:"records"`
	Total   int64       `json:"total"`
}
