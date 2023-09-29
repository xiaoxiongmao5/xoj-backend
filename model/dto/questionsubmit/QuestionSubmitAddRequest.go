/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-29 01:01:14
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-29 01:03:55
 * @FilePath: /xoj-backend/model/dto/questionsubmit/QuestionSubmitAddRequest.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package questionsubmit

// 创建请求
type QuestionSubmitAddRequest struct {
	// 编程语言
	Language string `json:"language"`
	// 用户代码
	Code string `json:"code"`
	// 题目 id
	QuestionId int64 `json:"questionId"`
}
