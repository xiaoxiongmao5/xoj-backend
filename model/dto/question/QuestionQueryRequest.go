/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-26 22:44:05
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-27 22:24:49
 * @FilePath: /xoj-backend/model/dto/question/questionAddRequest.go
 * @Description: 业务层面的封装
 */
package question

import "github.com/xiaoxiongmao5/xoj/xoj-backend/model/common"

type QuestionQueryRequest struct {
	common.PageRequest // 嵌套PageRequest结构体，相当于继承
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
	// 创建用户 id
	UserId int64 `json:"userId"`
}
