/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-30 12:16:11
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-30 12:18:21
 * @FilePath: /xoj-backend/service/userService.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package service

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/dto/user"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/entity"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/vo"
)

type UserService interface {
	// 用户注册
	func(ctx *context.Context, userAccount, userPassword, checkUserPassword string) (num int64)

	// 用户登录
	func(ctx *context.Context, userAccount, userPassword string) (loginUserVO vo.LoginUserVO)

	// 获取当前登录用户
	func(ctx *context.Context) *entity.User

	// 是否为管理员
	func(user *entity.User) bool

	// 用户注销
	func(ctx *context.Context) bool

	// 获取脱敏的已登录用户信息
	func(userObj *entity.User) vo.LoginUserVO

	// 获取脱敏的用户信息
	func(userObj *entity.User) vo.UserVO

	// 获取脱敏的用户信息列表
	func(list []*entity.User) (respdata []vo.UserVO)

	// 获取查询条件
	func(qs orm.QuerySeter, queryRequest user.UserQueryRequest) orm.QuerySeter

	func(ids []int64) ([]*entity.User, error)
	func(id int64) (*entity.User, error)
	func(userObj *entity.User) (int64, error)
	func(userObj *entity.User) error
	func(id int64) error
}
