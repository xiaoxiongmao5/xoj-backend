/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-10-02 13:27:42
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-02 16:03:03
 * @FilePath: /xoj-backend/judge/strategy/JudgeServiceImpl.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package judge

// import (
// 	"encoding/json"

// 	beego "github.com/beego/beego/v2/server/web"
// 	"github.com/xiaoxiongmao5/xoj/xoj-backend/judge/codesandbox"
// 	"github.com/xiaoxiongmao5/xoj/xoj-backend/judge/codesandbox/model"
// 	"github.com/xiaoxiongmao5/xoj/xoj-backend/judge/strategy"
// 	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/dto/question"
// 	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/entity"
// 	questionsubmitstatusenum "github.com/xiaoxiongmao5/xoj/xoj-backend/model/enums/QuestionSubmitStatusEnum"
// 	"github.com/xiaoxiongmao5/xoj/xoj-backend/myresq"
// 	questionservice "github.com/xiaoxiongmao5/xoj/xoj-backend/service/QuestionService"
// 	questionsubmitservice "github.com/xiaoxiongmao5/xoj/xoj-backend/service/QuestionSubmitService"
// )

// type JudgeServiceImpl struct {
// 	beego.Controller
// }

// func (this JudgeServiceImpl) DoJudge(questionsubmitId int64) *entity.QuestionSubmit {
// 	// 1）传入题目的提交 id，获取到对应的题目、提交信息（包含代码、编程语言等）
// 	questionsubmitObj, err := questionsubmitservice.GetById(questionsubmitId)
// 	if err != nil {
// 		myresq.Abort(this.Ctx, myresq.NOT_FOUND_ERROR, "提交信息不存在")
// 		return questionsubmitObj
// 	}
// 	questionObj, err := questionservice.GetById(questionsubmitObj.QuestionId)
// 	if err != nil {
// 		myresq.Abort(this.Ctx, myresq.NOT_FOUND_ERROR, "题目不存在")
// 		return questionsubmitObj
// 	}
// 	// 2）如果题目提交状态不为等待中，就不用重复执行了
// 	if questionsubmitObj.Status != questionsubmitstatusenum.WAITING.GetValue() {
// 		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "题目正在判题中")
// 		return questionsubmitObj
// 	}

// 	// 3）更改判题（题目提交）的状态为 “判题中”，防止重复执行
// 	questionsubmitObj.Status = questionsubmitstatusenum.WAITING.GetValue()
// 	err = questionsubmitservice.UpdateById(questionsubmitObj)
// 	if err != nil {
// 		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "题目状态更新错误")
// 		return questionsubmitObj
// 	}

// 	// 4）调用沙箱，获取到执行结果

// 	// 获取输入用例
// 	JudgeCaseStr := questionObj.JudgeCase
// 	var judgeCaseList []question.JudgeCase
// 	if err := json.Unmarshal([]byte(JudgeCaseStr), &judgeCaseList); err != nil {
// 		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "输入用例转换失败")
// 		return questionsubmitObj
// 	}
// 	inputList := make([]string, len(judgeCaseList))
// 	for i, v := range judgeCaseList {
// 		inputList[i] = v.Input
// 	}

// 	codesandboxImpl := codesandbox.CodeSandboxFactory("example")
// 	executeCodeResponse := codesandbox.CodeSandboxProxy{CodeSandbox: codesandboxImpl}.ExecuteCode(model.ExecuteCodeRequest{
// 		InputList: inputList,
// 		Code:      questionsubmitObj.Code,
// 		Language:  questionsubmitObj.Language,
// 	})

// 	// 5）根据沙箱的执行结果，设置题目的判题状态和信息
// 	judgeContext := strategy.JudgeContext{
// 		JudgeInfo:      executeCodeResponse.JudgeInfo,
// 		InputList:      inputList,
// 		OutputList:     executeCodeResponse.OutputList,
// 		JudgeCaseList:  judgeCaseList,
// 		Question:       *questionObj,
// 		QuestionSubmit: *questionsubmitObj,
// 	}
// 	judgeInfo := JudgeManager{}.DoJudge(judgeContext)

// 	// 6）修改数据库中的判题结果
// 	questionsubmitObj.Status = questionsubmitstatusenum.SUCCEED.GetValue()
// 	if s, err := json.Marshal(judgeInfo); err != nil {
// 		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "判题信息转换失败")
// 		return questionsubmitObj
// 	} else {
// 		questionsubmitObj.JudgeInfo = string(s)
// 	}
// 	if err := questionsubmitservice.UpdateById(questionsubmitObj); err != nil {
// 		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "题目状态更新错误")
// 		return questionsubmitObj
// 	}
// 	return questionsubmitObj

// }
