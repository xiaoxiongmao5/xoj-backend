/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-27 10:27:02
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-10 15:13:18
 */
package questionservice

import (
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
func ValidQuestion(ctx *context.Context, questionObj *entity.Question, add bool) {
	title := questionObj.Title
	content := questionObj.Content
	tags := questionObj.Tags
	answer := questionObj.Answer
	judgeCase := questionObj.JudgeCase
	judgeConfig := questionObj.JudgeConfig
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
func GetQuerySeter(qs orm.QuerySeter, queryRequest question.QuestionQueryRequest) orm.QuerySeter {
	id := queryRequest.Id
	title := queryRequest.Title
	content := queryRequest.Content
	tags := queryRequest.Tags
	answer := queryRequest.Answer
	userId := queryRequest.UserId
	sortField := queryRequest.PageRequest.SortField
	sortOrder := queryRequest.PageRequest.SortOrder

	// 构建查询条件
	if id != 0 {
		qs = qs.Filter("id", id) // WHERE id = id
	}
	if userId != 0 {
		qs = qs.Filter("userId", userId)
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
func GetQuestionVO(ctx *context.Context, questionObj *entity.Question) vo.QuestionVO {
	questionVO := vo.QuestionVO_Obj2Vo(questionObj)
	// 关联查询用户信息
	userObj, err := userservice.GetById(questionObj.UserId)
	if err != nil {
		mylog.Log.Errorf("查询userId=[%d]的用户信息失败, err=%v", questionObj.UserId, err.Error())
		questionVO.UserVO = vo.UserVO{}
		return questionVO
	}
	questionVO.UserVO = userservice.GetUserVO(userObj)
	return questionVO
}

// 获取脱敏的题目信息列表
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
		userID := questionObj.UserId
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
		userMap[user.Id] = user
	}

	// 填充用户信息
	for _, questionObj := range list {
		questionVO := vo.QuestionVO_Obj2Vo(questionObj)
		if userObj, exists := userMap[questionObj.UserId]; exists {
			questionVO.UserVO = userservice.GetUserVO(userObj)
		} else {
			mylog.Log.Errorf("查询userId=[%d]的用户信息失败", questionObj.UserId)
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

func UpdateById(questionObj *entity.Question, cols ...string) error {
	num, err := mydb.O.Update(questionObj, cols...)
	if err != nil {
		return err
	}
	if num == 0 {
		mylog.Log.Info("无更新影响条目")
		return nil
	}
	return nil
}

func RemoveById(id int64) error {
	questionObj, err := GetById(id)
	if err != nil {
		return nil
	}
	questionObj.IsDelete = 1
	num, err := mydb.O.Update(questionObj, "isDelete")
	if err != nil {
		return err
	}
	if num == 0 {
		return nil
	}
	return nil
}

func GetTotal() (int64, error) {
	num, err := mydb.O.QueryTable(new(entity.Question)).Filter("isDelete", 0).Count()
	if err != nil {
		mylog.Log.Errorf("Question 表 select count 出错, err=[%v]", err.Error())
	}
	return num, err
}
