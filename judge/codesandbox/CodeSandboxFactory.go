/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-10-02 12:14:47
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-02 15:11:37
 * @FilePath: /xoj-backend/judge/codesandbox/CodeSandbox.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package codesandbox

import "github.com/xiaoxiongmao5/xoj/xoj-backend/judge/codesandbox/impl"

// 代码沙箱工厂（根据字符串参数创建指定的代码沙箱实例）
func CodeSandboxFactory(codesandboxType string) CodeSandbox {
	switch codesandboxType {
	case "example":
		return impl.ExampleCodeSandbox{}
	case "remote":
		return impl.RemoteCodeSandbox{}
	case "thirdParty":
		return impl.ThirdPartyCodeSandbox{}
	default:
		return impl.ExampleCodeSandbox{}
	}
}
