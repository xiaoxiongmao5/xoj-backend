/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-10-10 17:11:45
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-10 20:08:25
 * @FilePath: /xoj-backend/middleware/authMiddleware.go
 * @Description: 捕获中断业务异常的中间件
 */
package middleware

import (
	"fmt"
	"time"

	"github.com/beego/beego/v2/server/web/context"
	"github.com/sirupsen/logrus"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mylog"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/myresq"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/utils"
)

type AppHook struct {
	RequestID string
}

func (h *AppHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *AppHook) Fire(entry *logrus.Entry) error {
	entry.Data["req_id"] = h.RequestID
	return nil
}

func LogMiddleware(ctx *context.Context) {
	startTime := time.Now()                 // 开始时间
	requestID := utils.CreateUniSessionId() // 1365038848
	ctx.Input.SetData("uniSessionId", requestID)

	// 添加 reqId 到每次请求的日志中
	h := &AppHook{RequestID: requestID}
	mylog.Log.AddHook(h)

	// 将标识添加到响应头中，以便客户端可以获取它
	ctx.Output.Header("X-Request-ID", requestID)

	// 记录请求信息
	domain, err := utils.GetDomainFromReferer(ctx.Input.Refer())
	if err != nil {
		mylog.Log.Error("获得请求来源域名 Referer 失败, err=: ", err.Error())
	}
	localIP, err := utils.GetLocalIP()
	if err != nil {
		mylog.Log.Error("获得本机IP失败, err=: ", err.Error())
	}

	mylog.Log.WithFields(logrus.Fields{
		"请求路径":   ctx.Input.URL(),
		"请求方式":   ctx.Input.Method(),
		"目标Host": ctx.Input.Domain(),
		"来源域名":   domain,
		"来源IP":   utils.GetRequestIp(ctx.Input.IP()),
		"本机IP":   localIP,
	}).Info("请求日志")

	ctx.Input.SetData("startTime", startTime)
}

func LogMiddlewareAfter(ctx *context.Context) {
	startTimeInterface := ctx.Input.GetData("startTime")
	startTime, ok := startTimeInterface.(time.Time)
	if !ok {
		myresq.Abort(ctx, myresq.GET_CONTEXT_ERROR, "")
		return
	}
	endTime := time.Now() // 结束时间
	respStatus := ctx.Output.Status
	latencyTm := endTime.Sub(startTime) // 执行时间总耗时
	totaltm := ""
	if latencyTm < 1*time.Millisecond {
		totaltm = fmt.Sprintf("%dµs", latencyTm.Microseconds()) // 微秒 1微秒 = 1000纳秒
	} else if latencyTm < 1*time.Second {
		totaltm = fmt.Sprintf("%dms", latencyTm.Milliseconds()) // 毫秒 1毫秒 = 1000微秒
	} else {
		totaltm = fmt.Sprintf("%.2fs", latencyTm.Seconds()) // 秒 1秒 = 1000毫秒
	}

	mylog.Log.WithFields(logrus.Fields{
		"响应码": respStatus,
		"总耗时": totaltm,
	}).Info("响应日志")
}
