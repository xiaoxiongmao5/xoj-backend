/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-10-10 17:11:45
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-10 17:51:15
 * @FilePath: /xoj-backend/middleware/authMiddleware.go
 * @Description: 判断已登录状态的中间件
 */
package middleware

import (
	"github.com/beego/beego/v2/server/web/context"
	"github.com/sirupsen/logrus"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mylog"
	userservice "github.com/xiaoxiongmao5/xoj/xoj-backend/service/userService"
)

func AuthMiddleware(ctx *context.Context) {
	loginUser := userservice.GetLoginUser(ctx)
	ctx.Input.SetData("loginUser", loginUser)

	mylog.Log.WithFields(logrus.Fields{
		"pass": true,
	}).Info("middleware-登录校验")
}
