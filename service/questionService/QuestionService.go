/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-27 10:27:02
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-29 21:30:18
 * @FilePath: /xoj-backend/service/question.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package questionservice

import (
	"errors"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/constant"
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
//	@param questionObj
//	@param add
func ValidQuestion(ctx *context.Context, questionObj *entity.Question, add bool) {
	title := questionObj.Title
	content := questionObj.Content
	tags := questionObj.Tags
	answer := questionObj.Answer
	judgeCase := questionObj.Judgecase
	judgeConfig := questionObj.Judgeconfig
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

// 获取查询条件（使用 beego 的 ORM 来构建数据库查询条件（用户根据哪些字段查询，根据前端传来的请求对象））
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
	qs = qs.Filter("isDelete", 0)
	return qs
}

// 获取题目封装
//
//	@param questionObj
//	@return vo.QuestionVO
func GetQuestionVO(ctx *context.Context, questionObj *entity.Question) vo.QuestionVO {
	questionVO := vo.QuestionVO_Obj2Vo(questionObj)
	// 关联查询用户信息
	userObj, err := userservice.GetById(questionObj.Userid)
	if err != nil {
		mylog.Log.Errorf("查询userId=[%d]的用户信息失败, err=%v", questionObj.Userid, err.Error())
		questionVO.UserVO = vo.UserVO{}
		return questionVO
	}
	questionVO.UserVO = userservice.GetUserVO(userObj)
	return questionVO
}

// 获取脱敏的题目信息列表
//
//	@param ctx
//	@param list
//	@return respdata
func ListQuestionVO(ctx *context.Context, list []*entity.Question) (respdata []vo.QuestionVO) {
	if utils.IsEmpty(list) {
		return
	}
	// 定义一个map，用于存储查询到的用户信息
	userMap := make(map[int64]*entity.User)

	// 获取 list 中的用户ID列表
	var userIds []int64
	userIdsMap := make(map[int64]bool)
	for _, questionObj := range list {
		userID := questionObj.Userid
		if !userIdsMap[userID] {
			userIdsMap[userID] = true
			userIds = append(userIds, userID)
		}
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
	for _, questionObj := range list {
		questionVO := vo.QuestionVO_Obj2Vo(questionObj)
		if userObj, exists := userMap[questionObj.Userid]; exists {
			questionVO.UserVO = userservice.GetUserVO(userObj)
		} else {
			mylog.Log.Errorf("查询userId=[%d]的用户信息失败", questionObj.Userid)
			questionVO.UserVO = vo.UserVO{}
		}
		respdata = append(respdata, questionVO)
	}
	return
}

func GetById(id int64) (*entity.Question, error) {
	var questionObj entity.Question
	err := mydb.O.QueryTable(new(entity.Question)).Filter("id", id).Filter("isDelete", 0).One(&questionObj)
	if err == orm.ErrMultiRows {
		mylog.Log.Errorf("Question 表中存在 id=[%d] 的多条记录, qs.One err=[%v]", id, err.Error())
		return nil, err
	}
	if err == orm.ErrNoRows {
		mylog.Log.Errorf("Question 表没有找到 id=[%d] 的记录, qs.One err=[%v]", id, err.Error())
		return nil, err
	}
	return &questionObj, nil
}

func Save(questionObj *entity.Question) (int64, error) {
	num, err := mydb.O.Insert(questionObj)
	if err != nil {
		return -1, err
	}
	return num, nil
}

func UpdateById(questionObj *entity.Question) error {
	num, err := mydb.O.Update(questionObj)
	if err != nil {
		return err
	}
	if num == 0 {
		return errors.New("无更新影响条目")
	}
	return nil
}

func RemoveById(id int64) error {
	questionObj, err := GetById(id)
	if err != nil {
		return nil
	}
	questionObj.IsDelete = 1
	num, err := mydb.O.Update(questionObj)
	if err != nil {
		return err
	}
	if num == 0 {
		return nil
	}
	return nil
}
