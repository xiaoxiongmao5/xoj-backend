/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-10-10 14:04:32
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-10 14:15:07
 * @FilePath: /xoj-backend/myredis/myredis.go
 */
package myredis

import (
	"github.com/go-redis/redis/v8"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/config"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mylog"
)

var RedisCli *redis.Client

func init() {
	mylog.Log.Info("init begin: myredis")

	// 创建 Redis 客户端连接
	RedisCli = redis.NewClient(&redis.Options{
		Addr:     config.AppConfig.Redis.Addr, //Redis 服务器地址
		Username: config.AppConfig.Redis.UserName,
		Password: config.AppConfig.Redis.PassWord,
		DB:       config.AppConfig.Redis.DB,
	})

	mylog.Log.Info("init end  : myredis")
}

func Close(redisCli *redis.Client) error {
	if redisCli != nil {
		return redisCli.Close()
	}
	return nil
}
