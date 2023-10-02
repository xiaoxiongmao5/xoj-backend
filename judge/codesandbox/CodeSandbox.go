/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-10-02 12:14:47
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-02 15:20:04
 * @FilePath: /xoj-backend/judge/codesandbox/CodeSandbox.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package codesandbox

import "github.com/xiaoxiongmao5/xoj/xoj-backend/judge/codesandbox/model"

// 代码沙箱
type CodeSandbox interface {
	ExecuteCode(executeCodeRequest model.ExecuteCodeRequest) model.ExecuteCodeResponse
}
