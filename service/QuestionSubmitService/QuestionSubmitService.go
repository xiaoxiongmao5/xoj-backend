/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-29 09:20:16
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-02 15:51:02
 */
package questionsubmitservice

import (
	"errors"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/constant"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/dto/questionsubmit"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/entity"
	questionsubmitstatusenum "github.com/xiaoxiongmao5/xoj/xoj-backend/model/enums/QuestionSubmitStatusEnum"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/vo"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mydb"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mylog"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/myresq"
	questionservice "github.com/xiaoxiongmao5/xoj/xoj-backend/service/QuestionService"
	userservice "github.com/xiaoxiongmao5/xoj/xoj-backend/service/userService"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/utils"
)

// 题目提交
func DoQuestionSubmit(ctx *context.Context, params questionsubmit.QuestionSubmitAddRequest, userObj *entity.User) int64 {
	// 校验编程语言是否合法
	// todo
	if params.Language != "go" {
		myresq.Abort(ctx, myresq.PARAMS_ERROR, "编程语言错误")
		return -1
	}

	// 判断实体是否存在，根据类别获取实体
	questionObj, err := questionservice.GetById(params.QuestionId)
	if err != nil {
		myresq.Abort(ctx, myresq.NOT_FOUND_ERROR, "题目未找到")
		return -1
	}

	// 是否已提交题目

	// 每个用户串行提交题目
	questionSubmitObj := entity.QuestionSubmit{
		UserId:     userObj.ID,
		QuestionId: questionObj.ID,
		Code:       params.Code,
		Language:   params.Language,
	}
	// 设置初始状态
	questionSubmitObj.Status = questionsubmitstatusenum.WAITING.GetValue()
	questionSubmitObj.JudgeInfo = "{}"

	id, err := Save(&questionSubmitObj)
	if err != nil {
		myresq.Abort(ctx, myresq.OPERATION_ERROR, "数据插入失败")
		return -1
	}

	// 执行判题服务
	// todo 异步执行
	// CompletableFuture.runAsync(() -> {
	// 	judgeService.doJudge(id);
	// });
	// judge.JudgeServiceImpl{}.DoJudge(id)

	return id
}

// 获取查询条件
func GetQuerySeter(qs orm.QuerySeter, queryRequest questionsubmit.QuestionSubmitQueryRequest) orm.QuerySeter {
	language := queryRequest.Language
	status := queryRequest.Status
	questionId := queryRequest.QuestionId
	userId := queryRequest.UserId
	sortField := queryRequest.PageRequest.SortField
	sortOrder := queryRequest.PageRequest.SortOrder

	if utils.IsNotBlank(language) {
		qs = qs.Filter("language", language)
	}
	if utils.IsNotBlank(userId) {
		qs = qs.Filter("userId", userId)
	}
	if utils.IsNotBlank(questionId) {
		qs = qs.Filter("questionId", questionId)
	}
	// todo
	// queryWrapper.eq(QuestionSubmitStatusEnum.getEnumByValue(status) != null, "status", status);
	if utils.IsNotBlank(status) {
		qs = qs.Filter("status", status)
	}

	if utils.IsNotBlank(sortField) {
		order := sortField
		if utils.CheckSame[string]("检查排序是否一样", sortOrder, constant.SORT_ORDER_DESC) {
			order = "-" + order
		}
		qs = qs.OrderBy(order)
	}
	qs = qs.Filter("isDelete", 0)
	return qs
}

// 获取提交题目的封装
func GetQuestionSubmitVO(ctx *context.Context, questionSubmitObj *entity.QuestionSubmit, loginUser *entity.User) vo.QuestionSubmitVO {
	questionSubmitVO := vo.QuestionSubmitVO_Obj2Vo(questionSubmitObj)
	// 脱敏：仅本人和管理员能看见自己（提交 userId 和登录用户 id 不同）提交的代码
	userId := loginUser.ID
	// 处理脱敏
	if !utils.CheckSame[int64]("检查当前用户与题目所属用户id是否一致", userId, questionSubmitObj.UserId) && !userservice.IsAdmin(loginUser) {
		questionSubmitVO.Code = ""
	}

	questionSubmitVO.UserVO = userservice.GetUserVO(loginUser)

	if questionObj, err := questionservice.GetById(questionSubmitVO.QuestionId); err != nil {
		mylog.Log.Errorf("查询questionId=[%d]的题目信息失败, err=%v", questionSubmitVO.QuestionId, err.Error())
		questionSubmitVO.QuestionVO = vo.QuestionVO{}
	} else {
		questionSubmitVO.QuestionVO = questionservice.GetQuestionVO(ctx, questionObj)
	}

	return questionSubmitVO
}

// 获取脱敏的提交题目信息列表
func ListQuestionSubmitVOPage(ctx *context.Context, list []*entity.QuestionSubmit, loginUser *entity.User) (respdata []vo.QuestionSubmitVO) {
	if utils.IsEmpty(list) {
		return
	}
	for _, one := range list {
		respdata = append(respdata, GetQuestionSubmitVO(ctx, one, loginUser))
	}
	return
}

func ListByIds(ids []int64) ([]*entity.QuestionSubmit, error) {
	qs := mydb.O.QueryTable(new(entity.QuestionSubmit))
	qs = qs.Filter("id__in", ids).Filter("isDelete", 0)
	var questionSubmits []*entity.QuestionSubmit
	_, err := qs.All(&questionSubmits)
	if err != nil {
		mylog.Log.Errorf("User ListByIds qs.All error: %v", err.Error())
		return questionSubmits, err
	}
	return questionSubmits, nil
}

func GetById(id int64) (*entity.QuestionSubmit, error) {
	var questionSubmitObj entity.QuestionSubmit
	err := mydb.O.QueryTable(new(entity.QuestionSubmit)).Filter("id", id).Filter("isDelete", 0).One(&questionSubmitObj)
	if err == orm.ErrMultiRows {
		mylog.Log.Errorf("QuestionSubmit 表中存在 id=[%d] 的多条记录, qs.One err=[%v]", id, err.Error())
		return nil, err
	}
	if err == orm.ErrNoRows {
		mylog.Log.Errorf("QuestionSubmit 表没有找到 id=[%d] 的记录, qs.One err=[%v]", id, err.Error())
		return nil, err
	}
	return &questionSubmitObj, nil
}

func Save(questionSubmitObj *entity.QuestionSubmit) (int64, error) {
	num, err := mydb.O.Insert(questionSubmitObj)
	if err != nil {
		return -1, err
	}
	return num, nil
}

func UpdateById(questionSubmitObj *entity.QuestionSubmit) error {
	num, err := mydb.O.Update(questionSubmitObj)
	if err != nil {
		return err
	}
	if num == 0 {
		return errors.New("无更新影响条目")
	}
	return nil
}

func RemoveById(id int64) error {
	questionSubmitObj, err := GetById(id)
	if err != nil {
		return nil
	}
	questionSubmitObj.IsDelete = 1
	num, err := mydb.O.Update(questionSubmitObj)
	if err != nil {
		return err
	}
	if num == 0 {
		return nil
	}
	return nil
}

func GetTotal() (int64, error) {
	num, err := mydb.O.QueryTable(new(entity.QuestionSubmit)).Filter("isDelete", 0).Count()
	if err != nil {
		mylog.Log.Errorf("QuestionSubmit 表 select count 出错, err=[%v]", err.Error())
	}
	return num, err
}
