package main

import (
	"net/http"

	"github.com/xiaoxiongmao5/xoj/xoj-backend/config"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/loadconfig"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mydb"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mylog"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/myresq"
	_ "github.com/xiaoxiongmao5/xoj/xoj-backend/routers"
	_ "github.com/xiaoxiongmao5/xoj/xoj-backend/store"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

//	@title			XOJ 项目
//	@version		1.0
//	@description	在线判题系统
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	小熊
//	@contact.url	https://github.com/xiaoxiongmao5
//	@contact.email	627516430@qq.com

//	@license.name	license.name
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host	localhost:8091
func main() {
	var err error

	// 实例化日志对象
	if mylog.Log, err = mylog.SetupLogger(); err != nil {
		panic(err)
	}
	defer mylog.Log.Writer().Close()

	// 加载App配置数据
	if config.AppConfig, err = loadconfig.LoadAppConfig(); err != nil {
		panic(err)
	}
	// 加载APP动态配置数据
	if config.AppConfigDynamic, err = loadconfig.LoadAppConfigDynamic(); err != nil {
		panic(err)
	}

	// 初始化数据库连接池
	if mydb.DB, err = mydb.ConnectionPool(config.AppConfig.Database.SavePath, config.AppConfig.Database.MaxOpenConns); err != nil {
		panic(err)
	}
	defer mydb.DB.Close()

	// 启动配置文件加载协程
	go loadconfig.LoadAppDynamicConfigCycle()

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

		// 开启监控：Admin 管理后台
		beego.BConfig.Listen.EnableAdmin = true
		// 修改监听的地址和端口：
		beego.BConfig.Listen.AdminAddr = "localhost"
		beego.BConfig.Listen.AdminPort = 8088
	}
	beego.BConfig.RecoverPanic = true
	beego.BConfig.RecoverFunc = func(ctx *context.Context, config *beego.Config) {
		if err := recover(); err != nil {
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

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	beego.Run()
}
