/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-26 22:44:05
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-27 17:41:51
 * @FilePath: /xoj-backend/model/dto/question/questionAddRequest.go
 * @Description: 业务层面的封装
 */
package question

import (
	"database/sql"
	"encoding/json"

	"github.com/beego/beego/v2/server/web/context"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/dbsq"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/myresq"
)

type QuestionAddRequest struct {
	//标题
	Title string `json:"title"`
	//内容
	Content string `json:"content"`
	//标签列表
	Tags []string `json:"tags"`
	//题目答案
	Answer string `json:"answer"`
	//判题用例
	Judgecase []JudgeCase `json:"judgecase"`
	//判题配置
	Judgeconfig JudgeConfig `json:"judgeconfig"`
}

func QuestionAddRequest2DBParams(ctx *context.Context, request *QuestionAddRequest) *dbsq.AddQuestionParams {
	ret := &dbsq.AddQuestionParams{
		Title:   sql.NullString{Valid: true, String: request.Title},
		Content: sql.NullString{Valid: true, String: request.Content},
		Answer:  sql.NullString{Valid: true, String: request.Answer},
	}
	if s, err := json.Marshal(request.Tags); err == nil && string(s) != "null" {
		ret.Tags = sql.NullString{Valid: true, String: string(s)}
	} else {
		myresq.Abort(ctx, myresq.PARAMS_ERROR, "")
		return nil
	}
	if s, err := json.Marshal(request.Judgecase); err == nil && string(s) != "null" {
		ret.Judgecase = sql.NullString{Valid: true, String: string(s)}
	} else {
		myresq.Abort(ctx, myresq.PARAMS_ERROR, "")
		return nil
	}
	if s, err := json.Marshal(request.Judgeconfig); err == nil && string(s) != "null" {
		ret.Judgeconfig = sql.NullString{Valid: true, String: string(s)}
	} else {
		myresq.Abort(ctx, myresq.PARAMS_ERROR, "")
		return nil
	}
	return ret
}
