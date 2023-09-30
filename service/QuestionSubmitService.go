/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-30 12:14:13
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-30 12:15:49
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

type QuestionSubmitService interface {
	// 题目提交
	func(ctx *context.Context, params questionsubmit.QuestionSubmitAddRequest, userObj *entity.User) int64

	// 获取查询条件
	func(qs orm.QuerySeter, queryRequest questionsubmit.QuestionSubmitQueryRequest) orm.QuerySeter

	// 获取提交题目的封装
	func(ctx *context.Context, questionSubmitObj *entity.QuestionSubmit, loginUser *entity.User) vo.QuestionSubmitVO

	// 获取脱敏的提交题目信息列表
	func(ctx *context.Context, list []*entity.QuestionSubmit, loginUser *entity.User) (respdata []vo.QuestionSubmitVO)

	func(ids []int64) ([]*entity.QuestionSubmit, error)
	func(id int64) (*entity.QuestionSubmit, error)
	func(questionSubmitObj *entity.QuestionSubmit) (int64, error)
	func(questionSubmitObj *entity.QuestionSubmit) error
	func(id int64) error
}
