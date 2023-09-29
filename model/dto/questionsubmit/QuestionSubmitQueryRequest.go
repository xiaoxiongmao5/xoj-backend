/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-29 01:01:24
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-29 01:09:31
 * @FilePath: /xoj-backend/model/dto/questionsubmit/QuestionSubmitQueryRequest.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package questionsubmit

import "github.com/xiaoxiongmao5/xoj/xoj-backend/model/common"

// 查询请求
type QuestionSubmitQueryRequest struct {
	common.PageRequest
	// 编程语言
	Language string `json:"language"`
	// 判题状态（0 - 待判题、1 - 判题中、2 - 成功、3 - 失败）
	Status int32 `json:"status"`
	// 题目 id
	QuestionId int64 `json:"questionId"`
	// 用户 id
	Userid int64 `json:"userId"`
}
