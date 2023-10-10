/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-30 11:37:36
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-10 22:09:42
 */
package swagtype

import (
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/entity"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/vo"
)

type QuestionResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    entity.Question `json:"data"`
}

type QuestionVOResponse struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    vo.QuestionVO `json:"data"`
}

type ListQuestionResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    listQuestion `json:"data"`
}
type listQuestion struct {
	Records []entity.Question `json:"records"`
	Total   int64             `json:"total"`
}

type ListQuestionVOResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    listQuestionVO `json:"data"`
}

type listQuestionVO struct {
	Records []vo.QuestionVO `json:"records"`
	Total   int64           `json:"total"`
}

type ListQuestionSubmitVOResponse struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    listQuestionSubmitVO `json:"data"`
}

type listQuestionSubmitVO struct {
	Records []vo.QuestionSubmitVO `json:"records"`
	Total   int64                 `json:"total"`
}

type QuestionSubmitVOResponse struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    vo.QuestionSubmitVO `json:"data"`
}
