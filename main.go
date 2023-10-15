/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-27 14:46:54
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-15 22:36:44
 * @FilePath: /xoj-backend/main.go
 */
package main

import (
	"github.com/xiaoxiongmao5/xoj/xoj-backend/middleware"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/myredis"
	_ "github.com/xiaoxiongmao5/xoj/xoj-backend/myrpc"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mysession"

	"github.com/xiaoxiongmao5/xoj/xoj-backend/config"
	_ "github.com/xiaoxiongmao5/xoj/xoj-backend/docs"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mylog"
	_ "github.com/xiaoxiongmao5/xoj/xoj-backend/routers"
	_ "github.com/xiaoxiongmao5/xoj/xoj-backend/store"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	mylog.Log.Info("init begin: main")

	mylog.Log.Info("init end  : main")
}

//	@title			XOJ 项目
//	@version		1.0
//	@description	在线判题系统
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	小熊
//	@contact.url	https://github.com/xiaoxiongmao5
//	@contact.email	627516430@qq.com

//	@license.name	license.name
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host	localhost:8091
func main() {
	defer mylog.Log.Writer().Close()
	defer myredis.Close(myredis.RedisCli)

	// 启动动态配置文件加载协程
	go config.LoadAppDynamicConfigCycle()

	if mysession.GlobalSessions != nil {
		go mysession.GlobalSessions.GC()
	} else {
		panic("GlobalSessions is nil, cannot start GC")
	}

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "docs"

		// 开启监控：Admin 管理后台
		beego.BConfig.Listen.EnableAdmin = true
		// 修改监听的地址和端口：
		beego.BConfig.Listen.AdminAddr = "localhost"
		beego.BConfig.Listen.AdminPort = 8088
	}

	// 全局异常捕获
	beego.BConfig.RecoverPanic = true
	beego.BConfig.RecoverFunc = middleware.ExceptionHandingMiddleware

	// // 处理跨域
	// beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
	// 	AllowOrigins:     []string{"http://localhost:8080", "https://*.jiexiong.com"}, //"*"
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// }))

	// // 使用session
	// // 设置是否开启 Session，默认是 false
	// beego.BConfig.WebConfig.Session.SessionOn = true
	// // 设置 Session 过期的时间，默认值是 3600 秒
	// beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3000

	// // 设置 Session 的引擎，默认是 memory，目前支持还有 file、mysql、redis 等
	// // beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	// // 设置对应 file、mysql、redis 引擎的保存路径或者链接地址，默认值是空
	// // beego.BConfig.WebConfig.Session.SessionProviderConfig =

	// // 设置 cookies 的名字，Session 默认是保存在用户的浏览器 cookies 里面的，默认名是 beegosessionID
	// beego.BConfig.WebConfig.Session.SessionName = "Test"

	// // 设置 cookie 的过期时间，cookie 是用来存储保存在客户端的数据
	// beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 3600

	// beego.BConfig.WebConfig.StaticDir = map[string]string{
	// 	"/swagger": "./docs",
	// }

	// beego.BConfig.WebConfig.StaticDir["/swagger"] = "docs"
	// beego.BConfig.WebConfig.StaticDir = map[string]string{
	// 	// prefix => directory
	// 	"/swagger": "./docs",
	// 	"/view":    "./views",
	// }

	beego.Run()
}
