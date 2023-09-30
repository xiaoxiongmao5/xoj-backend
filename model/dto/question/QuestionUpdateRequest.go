/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-26 22:44:05
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-30 10:36:22
 * @FilePath: /xoj-backend/model/dto/question/questionAddRequest.go
 * @Description: 业务层面的封装
 */
package question

type QuestionUpdateRequest struct {
	// id
	ID int64 `json:"id"`
	// 标题
	Title string `json:"title"`
	// 内容
	Content string `json:"content"`
	// 标签列表（json 数组）
	Tags []string `json:"tags"`
	// 题目答案
	Answer string `json:"answer"`
	// 判题用例
	JudgeCase []JudgeCase `json:"judgeCase"`
	// 判题配置
	JudgeConfig JudgeConfig `json:"judgeConfig"`
}
