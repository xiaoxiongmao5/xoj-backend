/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-27 10:27:02
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-28 23:31:30
 * @FilePath: /xoj-backend/service/question.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package questionservice

import (
	"database/sql"

	"context"

	"github.com/beego/beego/v2/client/orm"
	beecontext "github.com/beego/beego/v2/server/web/context"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/constant"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/dbsq"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/dto/question"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/entity"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/vo"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mydb"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mylog"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/myresq"
	userservice "github.com/xiaoxiongmao5/xoj/xoj-backend/service/userService"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/utils"
)

// 校验题目是否合法
//
//	@param ctx
//	@param question
//	@param add
func ValidQuestion(ctx *beecontext.Context, question *entity.Question, add bool) {
	title := question.Title
	content := question.Content
	tags := question.Tags
	answer := question.Answer
	judgeCase := question.Judgecase
	judgeConfig := question.Judgeconfig
	// 创建时，参数不能为空
	if add {
		if utils.IsAnyBlank(title, content, tags) {
			myresq.Abort(ctx, myresq.PARAMS_ERROR, "")
			return
		}
	}
	// 有参数则校验
	if len(title) > 80 {
		myresq.Abort(ctx, myresq.PARAMS_ERROR, "标题过长")
		return
	}
	if len(content) > 8192 {
		myresq.Abort(ctx, myresq.PARAMS_ERROR, "内容过长")
		return
	}
	if len(answer) > 8192 {
		myresq.Abort(ctx, myresq.PARAMS_ERROR, "答案过长")
		return
	}
	if len(judgeCase) > 8192 {
		myresq.Abort(ctx, myresq.PARAMS_ERROR, "判题用例过长")
		return
	}
	if len(judgeConfig) > 8192 {
		myresq.Abort(ctx, myresq.PARAMS_ERROR, "判题配置过长")
		return
	}
}

// 使用 beego 的 ORM 来构建数据库查询时分页条件
//
//	@param qs
//	@param current
//	@param pageSize
//	@return orm.QuerySeter
func GetQuerySeterByPage(qs orm.QuerySeter, current, pageSize int64) orm.QuerySeter {
	limit, offset := utils.CalculateLimitOffset[int64](current, pageSize)
	return qs.Limit(limit, offset)
}

// 使用 beego 的 ORM 来构建数据库查询条件（用户根据哪些字段查询，根据前端传来的请求对象）
//
// @param qs
// @param queryRequest
// @return orm.QuerySeter
func GetQuerySeter(qs orm.QuerySeter, queryRequest question.QuestionQueryRequest) orm.QuerySeter {
	id := queryRequest.ID
	title := queryRequest.Title
	content := queryRequest.Content
	tags := queryRequest.Tags
	answer := queryRequest.Answer
	userid := queryRequest.Userid
	sortField := queryRequest.PageRequest.SortField
	sortOrder := queryRequest.PageRequest.SortOrder

	// 构建查询条件
	if id != 0 {
		qs = qs.Filter("id", id) // WHERE id = id
	}
	if userid != 0 {
		qs = qs.Filter("userid", userid)
	}
	if utils.IsNotBlank(title) {
		qs = qs.Filter("title__icontains", title) // WHERE title LIKE '%title%'
	}
	if utils.IsNotBlank(content) {
		qs = qs.Filter("content__icontains", content)
	}
	if utils.IsNotBlank(answer) {
		qs = qs.Filter("answer__icontains", answer)
	}
	// 拼接标签查询条件
	if len(tags) > 0 {
		for _, tag := range tags {
			qs = qs.Filter("tags__icontains", tag)
		}
	}

	if utils.IsNotBlank(sortField) {
		order := sortField
		if utils.CheckSame[string]("检查排序是否一样", sortOrder, constant.SORT_ORDER_DESC) {
			order = "-" + order
		}
		qs = qs.OrderBy(order) // ORDER BY order DESC
	}
	qs = qs.Filter("isdelete", 0)
	return qs
}

// GetQuestionVO
//
//	@param original
//	@return vo.QuestionVO
func GetQuestionVO(ctx *beecontext.Context, original *entity.Question) vo.QuestionVO {
	questionVO := vo.Obj2Vo(original)
	// 关联查询用户信息
	userInfo, err := userservice.GetById(original.Userid)
	if err != nil {
		userInfo = &entity.User{}
		mylog.Log.Errorf("查询userId=[%d]的用户信息失败, err=%v", original.Userid, err.Error())
		myresq.Abort(ctx, myresq.OPERATION_ERROR, "查询失败")
		return questionVO
	}
	questionVO.UserVO = userservice.GetUserVO(userInfo)
	return questionVO
}

// GetQuestionVOPage
//
//	@param ctx
//	@param originalPage
//	@return respdata
func GetQuestionVOPage(ctx *beecontext.Context, originalPage []*entity.Question) (respdata []vo.QuestionVO) {
	if utils.IsEmpty(originalPage) {
		return
	}
	// 定义一个map，用于存储查询到的用户信息
	userMap := make(map[int64]*entity.User)

	// 获取 originalPage 中的用户ID列表
	var userIds []int64
	for _, questionObj := range originalPage {
		userIds = append(userIds, questionObj.Userid)
	}

	// 查询用户信息
	users, err := userservice.ListByIds(userIds)
	if err != nil {
		// 处理查询错误
		mylog.Log.Errorf("批量查询userIds=[%v]的用户信息失败[userservice.ListByIds], err=[%v]", userIds, err.Error())
		myresq.Abort(ctx, myresq.OPERATION_ERROR, "查询失败")
		return
	}

	// 将查询到的用户信息存储到userMap中
	for _, user := range users {
		userMap[user.ID] = user
	}

	// 填充用户信息
	for _, questionInfo := range originalPage {
		questionVO := vo.Obj2Vo(questionInfo)
		questionVO.UserVO = userservice.GetUserVO(userMap[questionInfo.Userid])
		respdata = append(respdata, questionVO)
	}
	return
}

func Save(dbparams *dbsq.AddQuestionParams) (sql.Result, error) {
	conn, err := mydb.GetConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	q := dbsq.New(conn)
	ctx := context.Background()
	return q.AddQuestion(ctx, dbparams)
}

func GetById(id int64) (*dbsq.Question, error) {
	conn, err := mydb.GetConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	q := dbsq.New(conn)
	ctx := context.Background()
	return q.GetQuestionById(ctx, id)
}

func UpdateById(dbparams *dbsq.UpdateQuestionParams) error {
	conn, err := mydb.GetConn()
	if err != nil {
		return err
	}
	defer conn.Close()
	q := dbsq.New(conn)
	ctx := context.Background()
	return q.UpdateQuestion(ctx, dbparams)
}

func RemoveById(id int64) error {
	conn, err := mydb.GetConn()
	if err != nil {
		return err
	}
	defer conn.Close()
	q := dbsq.New(conn)
	ctx := context.Background()
	return q.DeleteQuestion(ctx, id)
}
