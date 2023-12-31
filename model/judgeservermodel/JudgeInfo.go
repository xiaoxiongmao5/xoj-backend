/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-29 09:12:38
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-15 18:54:02
 * @FilePath: /xoj-backend/model/dto/questionsubmit/JudgeInfo.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package judgeservermodel

type JudgeInfo struct {
	// 程序执行信息(这是由判题系统写的)
	Message string `json:"message"`
	// 程序执行信息详细(这是由判题系统写的)
	Detail string `json:"detail"`
	// 消耗内存
	Memory int64 `json:"memory"`
	// 消耗时间
	Time int64 `json:"time"`
}
