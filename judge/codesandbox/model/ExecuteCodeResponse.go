/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-10-02 12:27:09
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-02 15:13:36
 * @FilePath: /xoj-backend/judge/codesandbox/model/ExecuteCodeResponse.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

type ExecuteCodeResponse struct {
	OutputList []string
	// 接口信息
	Message string `json:"message"`
	// 执行状态
	Status int32 `json:"status"`
	// 判题信息
	JudgeInfo JudgeInfo `json:"judgeInfo"`
}
