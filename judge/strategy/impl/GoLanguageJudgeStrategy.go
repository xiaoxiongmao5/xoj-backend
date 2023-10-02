/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-10-02 14:25:03
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-02 15:04:49
 * @FilePath: /xoj-backend/judge/strategy/impl/DefaultJudgeStrategy.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package impl

import (
	"encoding/json"

	"github.com/xiaoxiongmao5/xoj/xoj-backend/judge/codesandbox/model"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/judge/strategy"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/dto/question"
	judgeinfomessageenum "github.com/xiaoxiongmao5/xoj/xoj-backend/model/enums/JudgeInfoMessageEnum"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mylog"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/utils"
)

// Go 程序的判题策略
type GoLanguageJudgeStrategy struct {
}

// 执行判题
func (this GoLanguageJudgeStrategy) DoJudge(judgeContext strategy.JudgeContext) model.JudgeInfo {
	inputList := judgeContext.InputList
	outputList := judgeContext.OutputList
	judgeCaseList := judgeContext.JudgeCaseList
	quesionObj := judgeContext.Question
	judgeInfo := judgeContext.JudgeInfo
	memory := judgeInfo.Memory
	time := judgeInfo.Time

	judgeInfoResponse := model.JudgeInfo{
		Memory: memory,
		Time:   time,
	}

	// 先判断沙箱执行的结果输出数量是否和预期输出数量相等
	if !utils.CheckSame[int]("判断沙箱执行的输入和输出数量是否一致", len(inputList), len(outputList)) {
		judgeInfoResponse.Message = judgeinfomessageenum.WRONG_ANSWER.GetValue()
		return judgeInfoResponse
	}

	// 依次判断每一项输出和预期输出是否相等
	for i, len := 0, len(judgeCaseList); i < len; i++ {
		judgeCase := judgeCaseList[i]
		if !utils.CheckSame[string]("判断每项输出是否符合预期", judgeCase.Output, outputList[i]) {
			judgeInfoResponse.Message = judgeinfomessageenum.WRONG_ANSWER.GetValue()
			return judgeInfoResponse
		}
	}

	// 判断题目限制
	judgeConfigStr := quesionObj.JudgeConfig
	judgeConfig := question.JudgeConfig{}
	if err := json.Unmarshal([]byte(judgeConfigStr), &judgeConfig); err != nil {
		mylog.Log.Errorf("json.Unmarshal转换失败[%v]", judgeConfigStr)
		judgeInfoResponse.Message = judgeinfomessageenum.WRONG_ANSWER.GetValue()
		return judgeInfoResponse
	}
	needMemoryLimit := judgeConfig.MemoryLimit
	needTimeLimit := judgeConfig.TimeLimit
	if memory > needMemoryLimit {
		judgeInfoResponse.Message = judgeinfomessageenum.MEMORY_LIMIT_EXCEEDED.GetValue()
		return judgeInfoResponse
	}
	if time > needTimeLimit {
		judgeInfoResponse.Message = judgeinfomessageenum.TIME_LIMIT_EXCEEDED.GetValue()
		return judgeInfoResponse
	}
	judgeInfoResponse.Message = judgeinfomessageenum.ACCEPTED.GetValue()
	return judgeInfoResponse
}
