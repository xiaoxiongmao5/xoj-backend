/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-10-02 12:29:56
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-02 15:20:33
 * @FilePath: /xoj-backend/judge/codesandbox/impl/ExampleCodeSandbox.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package impl

import (
	"fmt"

	"github.com/xiaoxiongmao5/xoj/xoj-backend/judge/codesandbox/model"
)

type ThirdPartyCodeSandbox struct {
}

func (this ThirdPartyCodeSandbox) ExecuteCode(executeCodeRequest model.ExecuteCodeRequest) model.ExecuteCodeResponse {
	fmt.Println("第三方代码沙箱")
	return model.ExecuteCodeResponse{}
}
