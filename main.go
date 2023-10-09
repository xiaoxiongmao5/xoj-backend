/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-27 14:46:54
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-08 23:45:31
 * @FilePath: /xoj-backend/main.go
 */
package main

import (
	"net/http"

	_ "github.com/xiaoxiongmao5/xoj/xoj-backend/loadconfig"

	"github.com/xiaoxiongmao5/xoj/xoj-backend/config"
	_ "github.com/xiaoxiongmao5/xoj/xoj-backend/docs"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mylog"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/myresq"
	_ "github.com/xiaoxiongmao5/xoj/xoj-backend/routers"
	_ "github.com/xiaoxiongmao5/xoj/xoj-backend/store"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/beego/beego/v2/server/web/filter/cors"
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

	// 启动配置文件加载协程
	go config.LoadAppDynamicConfigCycle()

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
	beego.BConfig.RecoverFunc = func(ctx *context.Context, config *beego.Config) {
		if err := recover(); err != nil {
			mylog.Log.Errorf("beego.BConfig.RecoverFunc err= %v \n", err)

			// 从 Context 中获取错误码和消息
			response, ok := ctx.Input.GetData("json").(*myresq.BaseResponse)
			if !ok {
				response = myresq.NewBaseResponse(500, "未知错误", nil)
			}

			// 将 JSON 响应写入 Context，并设置响应头
			ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
			ctx.Output.SetStatus(http.StatusOK)
			ctx.Output.JSON(response, false, false)
		}
	}

	// 处理跨域
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"http://localhost:8080", "https://*.jiexiong.com"}, //"*"
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

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
