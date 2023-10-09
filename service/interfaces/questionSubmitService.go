/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-30 12:14:13
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-08 15:05:18
 * @FilePath: /xoj-backend/service/QuestionSubmitService.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package service

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/dto/questionsubmit"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/entity"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/vo"
)

type QuestionSubmitServiceInterface interface {
	// 题目提交
	DoQuestionSubmit(*context.Context, questionsubmit.QuestionSubmitAddRequest, *entity.User) int64

	// 获取查询条件
	GetQuerySeter(orm.QuerySeter, questionsubmit.QuestionSubmitQueryRequest) orm.QuerySeter

	// 获取提交题目的封装
	GetQuestionSubmitVO(*context.Context, *entity.QuestionSubmit, *entity.User) vo.QuestionSubmitVO

	// 获取脱敏的提交题目信息列表
	ListQuestionSubmitVOPage(*context.Context, []*entity.QuestionSubmit, *entity.User) []vo.QuestionSubmitVO

	ListByIds([]int64) ([]*entity.QuestionSubmit, error)
	GetById(int64) (*entity.QuestionSubmit, error)
	Save(*entity.QuestionSubmit) (int64, error)
	UpdateById(*entity.QuestionSubmit) error
	RemoveById(id int64) error
	GetTotal() (int64, error)
}
