/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-29 09:26:47
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-29 09:26:50
 * @FilePath: /xoj-backend/service/commonService/commonService.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package commonservice

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/utils"
)

// 获取分页查询条件（使用 beego 的 ORM 来构建数据库查询时分页条件）
//
//	@param qs
//	@param current
//	@param pageSize
//	@return orm.QuerySeter
func GetQuerySeterByPage(qs orm.QuerySeter, current, pageSize int64) orm.QuerySeter {
	limit, offset := utils.CalculateLimitOffset[int64](current, pageSize)
	return qs.Limit(limit, offset)
}
