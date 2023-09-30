/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-29 09:26:47
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-30 12:01:54
 */
package commonservice

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/utils"
)

// 获取分页查询条件（使用 beego 的 ORM 来构建数据库查询时分页条件）
func GetQuerySeterByPage(qs orm.QuerySeter, current, pageSize int64) orm.QuerySeter {
	limit, offset := utils.CalculateLimitOffset[int64](current, pageSize)
	return qs.Limit(limit, offset)
}
