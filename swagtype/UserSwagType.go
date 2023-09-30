/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-30 11:37:36
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-30 11:43:05
 * @FilePath: /xoj-backend/swagtype/UserSwagType.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
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
	Data  []entity.User `json:"data"`
	Total int64         `json:"total"`
}

type ListUserVOResponse struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    listUserVO `json:"data"`
}
type listUserVO struct {
	Data  []vo.UserVO `json:"data"`
	Total int64       `json:"total"`
}
