/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-10-10 14:26:59
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-10 16:06:20
 * @FilePath: /xoj-backend/produce/produce.go
 * @Description: 生产者
 */
package producer

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mylog"
)

const (
	QuestionSubmit2QueueKey = "Queue_QuestionSubmit"
)

// 将提交题目id推送到消息队列
func PushQuestionSubmit2Queue(ctx context.Context, client *redis.Client, id int64) error {
	if client == nil {
		return errors.New("redis client is nil")
	}
	err := client.RPush(ctx, QuestionSubmit2QueueKey, id).Err()
	if err != nil {
		mylog.Log.Error("[PushQuestionSubmit]无法将消息推送到队列, err=", err)
		return err
	}
	return nil
}
