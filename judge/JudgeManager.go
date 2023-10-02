/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-10-02 13:58:53
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-02 14:55:15
 * @FilePath: /xoj-backend/judge/strategy/JudgeManager.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package judge

import (
	"github.com/xiaoxiongmao5/xoj/xoj-backend/judge/codesandbox/model"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/judge/strategy"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/judge/strategy/impl"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/utils"
)

// 判题管理（简化调用）
type JudgeManager struct {
}

// 执行判题
func (this JudgeManager) DoJudge(judgeContext strategy.JudgeContext) model.JudgeInfo {
	var judgeStrategy strategy.JudgeStrategy = impl.DefaultJudgeStrategy{}

	questionSubmitObj := judgeContext.QuestionSubmit
	language := questionSubmitObj.Language

	if utils.CheckSame[string]("判断语言是否一致", language, "java") {
		judgeStrategy = impl.GoLanguageJudgeStrategy{}
	}

	return judgeStrategy.DoJudge(judgeContext)
}
