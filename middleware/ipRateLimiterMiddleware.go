/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-10-10 17:11:45
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-10 20:23:46
 * @FilePath: /xoj-backend/middleware/authMiddleware.go
 * @Description: 限流中间件
 */
package middleware

import (
	"errors"
	"sync"

	"github.com/beego/beego/v2/server/web/context"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/config"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mylog"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/myresq"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/utils"
	"golang.org/x/time/rate"
)

var (
	rateLimitMu sync.Mutex
	IPLimiter   *IPRateLimiter
)

func init() {
	mylog.Log.Info("init begin: middleware-IPRateLimiterMiddleware")

	// 创建IP限流器
	IPLimiter = NewIPRateLimiter()

	mylog.Log.Info("init end: middleware-IPRateLimiterMiddleware")
}

func NewIPRateLimiter() *IPRateLimiter {
	return &IPRateLimiter{
		limiter: make(map[string]*rate.Limiter),
	}
}

type IPRateLimiter struct {
	mu      sync.Mutex
	limiter map[string]*rate.Limiter
}

func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter, exists := i.limiter[ip]
	if !exists {
		limiter = rate.NewLimiter(rate.Limit(config.AppConfigDynamic.RateLimitConfig.RequestsPerSecond), config.AppConfigDynamic.RateLimitConfig.BucketSize)
		i.limiter[ip] = limiter
	}

	return limiter
}

// 定义一个中间件函数来进行限流
func IPRateLimiterMiddleware(ctx *context.Context) {
	ip := utils.GetRequestIp(ctx.Input.IP())

	if IPLimiter == nil {
		mylog.Log.Error("全局中间件 is nil")
	} else {
		limiter := IPLimiter.GetLimiter(ip)

		if !limiter.Allow() {
			myresq.Abort(ctx, myresq.TOO_MANY_REQUEST_ERROR, "")
			return
		}
	}
}

// 该方法用于动态更新具体IP的限流配置
func UpdateIPRateLimitConfig(ip string, requestsPerSecond float64, bucketSize int) error {
	// 在此处进行配置验证，确保新的配置是有效的
	if requestsPerSecond <= 0 || bucketSize <= 0 {
		return errors.New("Invalid configuration")
	}

	rateLimitMu.Lock()
	defer rateLimitMu.Unlock()

	limiter := IPLimiter.GetLimiter(ip)

	// 更新IP的限流配置
	limiter.SetLimit(rate.Limit(requestsPerSecond))
	limiter.SetBurst(bucketSize)

	return nil
}

// 该方法用于获得具体IP的限流配置
func GetIPRateLimitConfig(ip string) (requestsPerSecond float64, bucketSize int) {
	rateLimitMu.Lock()
	defer rateLimitMu.Unlock()

	limiter := IPLimiter.GetLimiter(ip)

	requestsPerSecond = float64(limiter.Limit())
	bucketSize = limiter.Burst()

	return
}
