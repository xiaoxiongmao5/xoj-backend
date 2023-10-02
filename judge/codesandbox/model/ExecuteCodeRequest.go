/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-10-02 12:24:44
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-02 12:26:53
 * @FilePath: /xoj-backend/judge/codesandbox/model/ExecuteCodeRequest.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

type ExecuteCodeRequest struct {
	InputList []string `json:"inputList"`
	Code      string   `json:"code"`
	Language  string   `json:"language"`
}
