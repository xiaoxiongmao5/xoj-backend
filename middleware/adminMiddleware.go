/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-10-10 17:11:45
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-10 18:04:34
 * @FilePath: /xoj-backend/middleware/authMiddleware.go
 * @Description: 判断是admin权限的中间件(需要判断已登录状态)
 */
package middleware

import (
	"github.com/beego/beego/v2/server/web/context"
	"github.com/sirupsen/logrus"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/entity"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mylog"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/myresq"
	userservice "github.com/xiaoxiongmao5/xoj/xoj-backend/service/userService"
)

func AdminMiddleware(ctx *context.Context) {
	// 从上下文中获取用户信息
	loginUserInterface := ctx.Input.GetData("loginUser")
	loginUser, ok := loginUserInterface.(*entity.User)
	if !ok {
		myresq.Abort(ctx, myresq.GET_CONTEXT_ERROR, "")
		return
	}

	// 判断当前用户是否是管理员
	if !userservice.IsAdmin(loginUser) {
		myresq.Abort(ctx, myresq.FORBIDDEN_ERROR, "")
		return
	}

	mylog.Log.WithFields(logrus.Fields{
		"pass": true,
	}).Info("middleware-Admin权限校验")
}
