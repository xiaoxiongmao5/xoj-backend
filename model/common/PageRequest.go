/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-27 15:14:58
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-27 22:23:27
 * @FilePath: /xoj-backend/model/idRequest.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package common

type PageRequest struct {
	// 当前页号
	Current int64 `json:"current"`
	// 页面大小
	PageSize int64 `json:"pageSize"`
	// 排序字段
	SortField string `json:"sortField"`
	// 排序顺序（默认升序）
	SortOrder string `json:"sortOrder"`
}
