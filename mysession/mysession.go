/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-10-11 15:21:10
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-12 14:50:54
 * @FilePath: /xoj-backend/mysession/mysession.go
 * @Description: mysession
 */
package mysession

import (
	"github.com/beego/beego/v2/server/web/session"
	// _ "github.com/beego/beego/v2/server/web/session/redis"

	"github.com/xiaoxiongmao5/xoj/xoj-backend/mylog"
)

var GlobalSessions *session.Manager

func init() {
	mylog.Log.Info("init begin: mysession")

	sessionConfig := &session.ManagerConfig{
		CookieName:      "xojSessionid", //客户端存储cookie的名字
		EnableSetCookie: true,           //是否开启 SetCookie, omitempty这个设置
		Gclifetime:      2592000,        //触发 GC 的时间, 设置和Maxlifetime一样(在服务端一到期时就立即被GC垃圾回收了)或者小于它
		Maxlifetime:     2592000,        //服务器端存储的数据的过期时间(秒 30天)
		Secure:          false,          //是否开启HTTPS，在cookie中设置cookie.Secure
		CookieLifeTime:  2592000,        //客户端存储的 cookie 的时间，默认值是 0，即浏览器生命周期
		ProviderConfig:  "./tmp",        //配置信息，根据不同的引擎设置不同的配置信息
		SessionIDPrefix: "xoj_",
	}
	GlobalSessions, _ = session.NewManager("memory", sessionConfig) //引擎名字，可以是memory、file、MySQL或redis(只有redis需要import _ "github.com/beego/beego/v2/server/web/session/redis")

	// sessionConfig := &session.ManagerConfig{
	// 	CookieName:      "xojSessionid",              //客户端存储cookie的名字
	// 	EnableSetCookie: true,                        //是否开启 SetCookie, omitempty这个设置
	// 	Gclifetime:      2592000,                     //触发 GC 的时间, 设置和Maxlifetime一样(在服务端一到期时就立即被GC垃圾回收了)或者小于它
	// 	Maxlifetime:     2592000,                     //服务器端存储的数据的过期时间(秒 30天)
	// 	Secure:          false,                       //是否开启HTTPS，在cookie中设置cookie.Secure
	// 	CookieLifeTime:  2592000,                     //客户端存储的 cookie 的时间，默认值是 0，即浏览器生命周期
	// 	ProviderConfig:  config.AppConfig.Redis.Addr, //配置信息，根据不同的引擎设置不同的配置信息，详细的配置请看下面的引擎设置(这里是：redis 的链接地址)
	// 	SessionIDPrefix: "xoj_",
	// }
	// GlobalSessions, _ = session.NewManager("redis", sessionConfig) //引擎名字，可以是memory、file、MySQL或redis(只有redis需要import _ "github.com/beego/beego/v2/server/web/session/redis")

	// go GlobalSessions.GC()

	mylog.Log.Info("init end: mysession")
}
