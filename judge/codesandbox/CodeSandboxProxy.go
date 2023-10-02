package codesandbox

import (
	"github.com/xiaoxiongmao5/xoj/xoj-backend/judge/codesandbox/model"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mylog"
)

type CodeSandboxProxy struct {
	CodeSandbox CodeSandbox
}

func (this CodeSandboxProxy) ExecuteCode(executeCodeRequest model.ExecuteCodeRequest) model.ExecuteCodeResponse {
	mylog.Log.Infof("代码沙箱请求信息：%v", executeCodeRequest)
	executeCodeResponse := this.CodeSandbox.ExecuteCode(executeCodeRequest)
	mylog.Log.Infof("代码沙箱请响应信息：%v", executeCodeResponse)
	return executeCodeResponse
}
