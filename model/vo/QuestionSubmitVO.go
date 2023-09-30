/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-27 10:39:07
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-30 10:23:59
 * @FilePath: /xoj-backend/model/vo/UserVO.go
 * @Description: 专门返回给前端用的，可以节约网络传输大小，或者过滤字段（脱敏）、保证安全性。
 */
package vo

import (
	"encoding/json"
	"time"

	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/dto/questionsubmit"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/entity"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mylog"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/utils"
)

type QuestionSubmitVO struct {
	// id
	ID int64 `json:"id"`
	// 编程语言
	Language string `json:"language"`
	// 用户代码
	Code      string                   `json:"code"`
	JudgeInfo questionsubmit.JudgeInfo `json:"judgeInfo"`
	// 判题状态（0 - 待判题、1 - 判题中、2 - 成功、3 - 失败）
	Status int32 `json:"status"`
	// 题目 id
	QuestionId int64 `json:"questionId"`
	// 用户 id
	UserId int64 `json:"userId"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 创建题目人的信息
	UserVO UserVO `json:"userVO"`
	// 对应题目信息
	QuestionVO QuestionVO `json:"questionVO"`
}

func QuestionSubmitVO_Vo2Obj(original *QuestionSubmitVO) entity.QuestionSubmit {
	data := entity.QuestionSubmit{}
	utils.CopyStructFields(*original, &data)
	if s, err := json.Marshal(original.JudgeInfo); err == nil && string(s) != "null" {
		data.JudgeInfo = string(s)
	}
	return data
}

func QuestionSubmitVO_Obj2Vo(original *entity.QuestionSubmit) QuestionSubmitVO {
	data := QuestionSubmitVO{}
	utils.CopyStructFields(*original, &data)
	if err := json.Unmarshal([]byte(original.JudgeInfo), &data.JudgeInfo); err != nil {
		mylog.Log.Warn("entity.QuestionSubmit 2 QuestionSubmitVO, json反序列化[JudgeInfo]错误", err.Error())
	}
	return data
}
