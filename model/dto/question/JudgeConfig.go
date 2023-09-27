/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-26 22:44:05
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-27 10:45:21
 * @FilePath: /xoj-backend/model/dto/question/questionAddRequest.go
 * @Description: 业务层面的封装
 */
package question

type JudgeConfig struct {
	//时间限制（ms）
	TimeLimit int64 `json:"timeLimit"`
	//内存限制（KB）
	MemoryLimit int64 `json:"memoryLimit"`
	//堆栈限制（KB）
	StackLimit int64 `json:"stackLimit"`
}
