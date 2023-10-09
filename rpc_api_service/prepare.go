/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-10-08 18:43:39
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-09 10:31:45
 */
package rpcapiservice

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertTimestampToTime(ts *timestamppb.Timestamp) time.Time {
	return ts.AsTime()
}
func ConvertTimeToTimestamp(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}
