/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-27 10:31:11
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-27 18:21:33
 * @FilePath: /xoj-backend/model/entity/Question.go
 * @Description: 实体类，表字段
 */
package entity

import (
	"time"

	"github.com/xiaoxiongmao5/xoj/xoj-backend/dbsq"
)

// 题目
type Question struct {
	// id
	ID int64 `json:"id"`
	// 标题
	Title string `json:"title"`
	// 内容
	Content string `json:"content"`
	// 标签列表（json 数组）
	Tags string `json:"tags"`
	// 题目答案
	Answer string `json:"answer"`
	// 题目提交数
	Submitnum int32 `json:"submitnum"`
	// 题目通过数
	Acceptednum int32 `json:"acceptednum"`
	// 判题用例（json 数组）
	Judgecase string `json:"judgecase"`
	// 判题配置（json 对象）
	Judgeconfig string `json:"judgeconfig"`
	// 点赞数
	Thumbnum int32 `json:"thumbnum"`
	// 收藏数
	Favournum int32 `json:"favournum"`
	// 创建用户 id
	Userid int64 `json:"userid"`
	// 创建时间
	Createtime time.Time `json:"createtime"`
	// 更新时间
	Updatetime time.Time `json:"updatetime"`
	// 是否删除
	Isdelete int32 `json:"isdelete"`
}

func DbConvertQuestion(original *dbsq.Question) *Question {
	converted := &Question{
		ID:          original.ID,
		Submitnum:   original.Submitnum,
		Acceptednum: original.Acceptednum,
		Thumbnum:    original.Thumbnum,
		Favournum:   original.Favournum,
		Userid:      original.Userid,
		Createtime:  original.Createtime,
		Updatetime:  original.Updatetime,
		Isdelete:    original.Isdelete,
	}

	if original.Title.Valid {
		converted.Title = original.Title.String
	}
	if original.Content.Valid {
		converted.Content = original.Content.String
	}
	if original.Tags.Valid {
		converted.Tags = original.Tags.String
	}
	if original.Answer.Valid {
		converted.Answer = original.Answer.String
	}
	if original.Judgecase.Valid {
		converted.Judgecase = original.Judgecase.String
	}
	if original.Judgeconfig.Valid {
		converted.Judgeconfig = original.Judgeconfig.String
	}

	return converted
}
