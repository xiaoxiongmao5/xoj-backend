/*
* @Author: 小熊 627516430@qq.com
* @Date: 2023-09-27 10:33:52
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-30 12:00:59
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
	Id int64 `json:"id"`
	// 标题
	Title string `json:"title"`
	// 内容
	Content string `json:"content"`
	// 标签列表（json 数组）
	Tags []string `json:"tags"`
	// 题目提交数
	SubmitNum int32 `json:"submitNum"`
	// 题目通过数
	AcceptedNum int32 `json:"acceptedNum"`
	// 判题配置（json 对象）
	JudgeConfig question.JudgeConfig `json:"judgeConfig"`
	// 点赞数
	ThumbNum int32 `json:"thumbNum"`
	// 收藏数
	FavourNum int32 `json:"favourNum"`
	// 创建用户 id
	UserId int64 `json:"userId"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 创建题目人的信息
	UserVO UserVO `json:"userVO"`
}

// 包装类转实体类
func QuestionVO_Vo2Obj(original *QuestionVO) entity.Question {
	data := entity.Question{}
	utils.CopyStructFields(original, &data)
	if s, err := json.Marshal(original.Tags); err == nil && string(s) != "null" {
		data.Tags = string(s)
	}
	if s, err := json.Marshal(original.JudgeConfig); err == nil && string(s) != "null" {
		data.JudgeConfig = string(s)
	}
	return data
}

// 实体类转包装类
func QuestionVO_Obj2Vo(original *entity.Question) QuestionVO {
	data := QuestionVO{}
	utils.CopyStructFields(*original, &data)
	if err := json.Unmarshal([]byte(original.Tags), &data.Tags); err != nil {
		mylog.Log.Warn("entity.Question 2 QuestionVO, json反序列化[Tags]错误", err.Error())
	}
	if err := json.Unmarshal([]byte(original.JudgeConfig), &data.JudgeConfig); err != nil {
		mylog.Log.Warn("entity.Question 2 QuestionVO, json反序列化[JudgeConfig]错误", err.Error())
	}
	return data
}
