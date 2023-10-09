/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-30 12:16:11
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-08 15:03:28
 */
package service

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/dto/user"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/entity"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/vo"
)

type UserServiceInterface interface {
	// 用户注册
	UserRegister(*context.Context, string, string, string) int64

	// 用户登录
	UserLogin(*context.Context, string, string) vo.LoginUserVO

	// 获取当前登录用户
	GetLoginUser(*context.Context) *entity.User

	// 是否为管理员
	IsAdmin(*entity.User) bool

	// 用户注销
	UserLogout(*context.Context) bool

	// 获取脱敏的已登录用户信息
	GetLoginUserVO(*entity.User) vo.LoginUserVO

	// 获取脱敏的用户信息
	GetUserVO(*entity.User) vo.UserVO

	// 获取脱敏的用户信息列表
	ListUserVO([]*entity.User) []vo.UserVO

	// 获取查询条件
	GetQuerySeter(orm.QuerySeter, user.UserQueryRequest) orm.QuerySeter

	ListByIds([]int64) ([]*entity.User, error)
	GetById(int64) (*entity.User, error)
	Save(*entity.User) (int64, error)
	UpdateById(*entity.User) error
	RemoveById(int64) error
	GetTotal() (int64, error)
}
