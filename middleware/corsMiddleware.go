/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-10-10 17:11:45
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-12 14:05:16
 * @FilePath: /xoj-backend/middleware/authMiddleware.go
 * @Description: 处理跨域请求的中间件
 */
package middleware

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	"github.com/sirupsen/logrus"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mylog"
)

func CORSMiddleware() web.HandleFunc {
	corsMiddleware := cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"}, //"*"
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})

	mylog.Log.WithFields(logrus.Fields{
		"pass": true,
	}).Info("middleware-处理跨域")

	return corsMiddleware
}
