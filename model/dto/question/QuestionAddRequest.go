/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-26 22:44:05
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-26 23:00:58
 * @FilePath: /xoj-backend/model/dto/question/questionAddRequest.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package question

type QuestionAddRequest struct {
	//标题
	Title string `json:"title"`
	//内容
	Content string `json:"content"`
	//标签列表
	Tags string `json:"tags"`
	//题目答案
	Answer string `json:"answer"`
	//判题用例
	Judgecase JudgeCase `json:"judgecase"`
	//判题配置
	Judgeconfig JudgeConfig `json:"judgeconfig"`
}
