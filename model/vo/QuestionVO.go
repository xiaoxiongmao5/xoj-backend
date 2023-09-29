/*
* @Author: 小熊 627516430@qq.com
* @Date: 2023-09-27 10:33:52
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-29 09:18:38
* @FilePath: /xoj-backend/model/vo/QuestionVO.go
* @Description: 专门返回给前端用的，可以节约网络传输大小，或者过滤字段（脱敏）、保证安全性。
*/
package vo

import (
	"encoding/json"
	"time"

	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/dto/question"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/entity"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mylog"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/utils"
)

type QuestionVO struct {
	// id
	ID int64 `json:"id"`
	// 标题
	Title string `json:"title"`
	// 内容
	Content string `json:"content"`
	// 标签列表（json 数组）
	Tags []string `json:"tags"`
	// 题目提交数
	Submitnum int32 `json:"submitnum"`
	// 题目通过数
	Acceptednum int32 `json:"acceptednum"`
	// 判题配置（json 对象）
	Judgeconfig question.JudgeConfig `json:"judgeconfig"`
	// 点赞数
	Thumbnum int32 `json:"thumbnum"`
	// 收藏数
	Favournum int32 `json:"favournum"`
	// 创建用户 id
	Userid int64 `json:"userid"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 创建题目人的信息
	UserVO UserVO `json:"uservo"`
}

// 包装类转实体类
//
//	@param original
//	@return *entity.Question
func QuestionVO_Vo2Obj(original *QuestionVO) entity.Question {
	data := entity.Question{}
	utils.CopyStructFields(original, &data)
	if s, err := json.Marshal(original.Tags); err == nil && string(s) != "null" {
		data.Tags = string(s)
	}
	if s, err := json.Marshal(original.Judgeconfig); err == nil && string(s) != "null" {
		data.Judgeconfig = string(s)
	}
	return data
}

// 实体类转包装类
//
//	@param original
//	@return *QuestionVO
func QuestionVO_Obj2Vo(original *entity.Question) QuestionVO {
	data := QuestionVO{}
	utils.CopyStructFields(*original, &data)
	if err := json.Unmarshal([]byte(original.Tags), &data.Tags); err != nil {
		mylog.Log.Warn("entity.Question 2 QuestionVO, json反序列化[Tags]错误", err.Error())
	}
	if err := json.Unmarshal([]byte(original.Judgeconfig), &data.Judgeconfig); err != nil {
		mylog.Log.Warn("entity.Question 2 QuestionVO, json反序列化[Judgeconfig]错误", err.Error())
	}
	return data
}
