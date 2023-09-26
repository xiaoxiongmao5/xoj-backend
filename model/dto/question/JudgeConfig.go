/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-26 22:44:05
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-26 22:59:57
 * @FilePath: /xoj-backend/model/dto/question/questionAddRequest.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
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
