/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-10-08 19:31:19
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-09 10:34:14
 */
package loadconfig

import (
	"flag"
	"os"

	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mylog"

	rpcapiservice "github.com/xiaoxiongmao5/xoj/xoj-backend/rpc_api_service"
)

func init() {
	mylog.Log.Info("init begin: loadrpc")

	config.SetProviderService(&rpcapiservice.QuestionServerImpl{})
	config.SetProviderService(&rpcapiservice.QuestionSubmitServerImpl{})

	// 加载 Dubbo-go 的配置
	LoadDubboConfig()

	mylog.Log.Info("init end  : loadrpc")
}

// 设置环境变量
func SetOsEnv() {
	// 使用命令行参数来指定配置文件路径
	configFile := flag.String("config", "conf/dubbogo.yaml", "Path to Dubbo-go config file")
	flag.Parse()

	// 设置 DUBBO_GO_CONFIG_PATH 环境变量
	os.Setenv("DUBBO_GO_CONFIG_PATH", *configFile)
}

// 加载 Dubbo-go 的配置
func LoadDubboConfig() {
	SetOsEnv()
	// 加载 Dubbo-go 的配置文件，根据环境变量 DUBBO_GO_CONFIG_PATH 中指定的配置文件路径加载配置信息。配置文件通常包括 Dubbo 服务的注册中心地址、协议、端口等信息。
	if err := config.Load(); err != nil {
		panic(err)
	}
}
