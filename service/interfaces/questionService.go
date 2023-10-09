/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-30 12:10:16
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-08 15:05:10
 * @FilePath: /xoj-backend/service/QuestionService.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package service

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/dto/question"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/entity"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/vo"
)

type QuestionServiceInterface interface {
	// 校验题目是否合法
	ValidQuestion(*context.Context, *entity.Question, bool)

	// 获取查询条件（使用 beego 的 ORM 来构建数据库查询条件（用户根据哪些字段查询，根据前端传来的请求对象））
	GetQuerySeter(orm.QuerySeter, question.QuestionQueryRequest) orm.QuerySeter

	// 获取题目封装
	GetQuestionVO(*context.Context, *entity.Question) vo.QuestionVO

	// 获取脱敏的题目信息列表
	ListQuestionVO(*context.Context, []*entity.Question) []vo.QuestionVO

	GetById(int64) (*entity.Question, error)
	Save(*entity.Question) (int64, error)
	UpdateById(*entity.Question) error
	RemoveById(int64) error
	GetTotal() (int64, error)
}
