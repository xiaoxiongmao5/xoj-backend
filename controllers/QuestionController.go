/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-26 22:18:34
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-10 22:13:05
 */
package controllers

import (
	"encoding/json"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/common"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/dto/question"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/dto/questionsubmit"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/entity"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mydb"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mylog"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/myresq"
	commonservice "github.com/xiaoxiongmao5/xoj/xoj-backend/service/commonService"
	questionservice "github.com/xiaoxiongmao5/xoj/xoj-backend/service/questionService"
	questionsubmitservice "github.com/xiaoxiongmao5/xoj/xoj-backend/service/questionSubmitService"
	userservice "github.com/xiaoxiongmao5/xoj/xoj-backend/service/userService"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/utils"
)

type QuestionController struct {
	beego.Controller
}

//	@Summary		添加
//	@Description	添加
//	@Tags			题目增删改查（管理员）
//	@Accept			application/json
//	@Produce		application/json
//	@Param			request	body		question.QuestionAddRequest	true	"添加参数"
//	@Success		200		{object}	object						"响应数据"
//	@Router			/question/add [post]
func (this QuestionController) AddQuestion() {
	var params question.QuestionAddRequest
	if err := this.BindJSON(&params); err != nil {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}

	questionObj := entity.Question{}
	if res := utils.CopyStructFields(params, &questionObj); !res {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}
	if s, err := json.Marshal(params.Tags); err == nil && string(s) != "null" {
		questionObj.Tags = string(s)
	}
	if s, err := json.Marshal(params.JudgeCase); err == nil && string(s) != "null" {
		questionObj.JudgeCase = string(s)
	}
	if s, err := json.Marshal(params.JudgeConfig); err == nil && string(s) != "null" {
		questionObj.JudgeConfig = string(s)
	}

	// 参数校验
	questionservice.ValidQuestion(this.Ctx, &questionObj, true)

	loginUserInterface := this.Ctx.Input.GetData("loginUser")
	loginUser, ok := loginUserInterface.(*entity.User)
	if !ok {
		myresq.Abort(this.Ctx, myresq.GET_CONTEXT_ERROR, "")
		return
	}
	questionObj.UserId = loginUser.Id

	id, err := questionservice.Save(&questionObj)
	if err != nil {
		mylog.Log.Error("添加题目失败, err=", err.Error())
		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "添加失败")
		return
	}

	// respdata := make(map[string]int64)
	// respdata["id"] = id
	myresq.Success(this.Ctx, id)
}

//	@Summary		删除
//	@Description	删除
//	@Tags			题目增删改查（管理员）
//	@Accept			application/json
//	@Produce		application/json
//	@Param			request	body		common.IdRequest	true	"删除id"
//	@Success		200		{object}	object				"响应数据"
//	@Router			/question/delete [post]
func (this QuestionController) DeleteQuestion() {
	var params common.IdRequest
	if err := this.BindJSON(&params); err != nil {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}

	loginUserInterface := this.Ctx.Input.GetData("loginUser")
	loginUser, ok := loginUserInterface.(*entity.User)
	if !ok {
		myresq.Abort(this.Ctx, myresq.GET_CONTEXT_ERROR, "")
		return
	}

	// 判断是否存在
	questionInfo, err := questionservice.GetById(params.Id)
	if err != nil || questionInfo.Id <= 0 {
		myresq.Abort(this.Ctx, myresq.NOT_FOUND_ERROR, "题目未找到")
		return
	}
	// 仅本人或管理员可删除
	if !utils.CheckSame[int64]("检查当前用户与题目所属用户id是否一致", questionInfo.UserId, loginUser.Id) && !userservice.IsAdmin(loginUser) {
		myresq.Abort(this.Ctx, myresq.NO_AUTH_ERROR, "")
		return
	}

	if err = questionservice.RemoveById(params.Id); err != nil {
		mylog.Log.Error("删除题目失败, err=", err.Error())
		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "删除失败")
	}

	myresq.Success(this.Ctx, nil)
}

//	@Summary		更新
//	@Description	更新
//	@Tags			用户增删改查
//	@Accept			application/json
//	@Produce		application/json
//	@Param			request	body		question.QuestionEditRequest	true	"更新参数"
//	@Success		200		{object}	object							"响应数据"
//	@Router			/question/edit [post]
func (this QuestionController) EditQuestion() {
	var params question.QuestionEditRequest
	if err := this.BindJSON(&params); err != nil {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}

	// 判断是否存在
	questionObj, err := questionservice.GetById(params.Id)
	if err != nil || questionObj.Id <= 0 {
		myresq.Abort(this.Ctx, myresq.NOT_FOUND_ERROR, "题目未找到")
		return
	}

	if res := utils.CopyStructFields(params, questionObj); !res {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}
	if s, err := json.Marshal(params.Tags); err == nil && string(s) != "null" {
		questionObj.Tags = string(s)
	}
	if s, err := json.Marshal(params.JudgeCase); err == nil && string(s) != "null" {
		questionObj.JudgeCase = string(s)
	}
	if s, err := json.Marshal(params.JudgeConfig); err == nil && string(s) != "null" {
		questionObj.JudgeConfig = string(s)
	}

	// 参数校验
	questionservice.ValidQuestion(this.Ctx, questionObj, false)

	loginUserInterface := this.Ctx.Input.GetData("loginUser")
	loginUser, ok := loginUserInterface.(*entity.User)
	if !ok {
		myresq.Abort(this.Ctx, myresq.GET_CONTEXT_ERROR, "")
		return
	}

	// 仅本人或管理员可编辑
	if !utils.CheckSame[int64]("检查当前用户与题目所属用户id是否一致", questionObj.UserId, loginUser.Id) && !userservice.IsAdmin(loginUser) {
		myresq.Abort(this.Ctx, myresq.NO_AUTH_ERROR, "")
		return
	}

	if err := questionservice.UpdateById(questionObj); err != nil {
		mylog.Log.Error("更新题目失败, err=", err.Error())
		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "更新失败")
		return
	}

	myresq.Success(this.Ctx, nil)
}

//	@Summary		更新
//	@Description	更新
//	@Tags			题目增删改查（管理员）
//	@Accept			application/json
//	@Produce		application/json
//	@Param			request	body		question.QuestionUpdateRequest	true	"更新参数"
//	@Success		200		{object}	object							"响应数据"
//	@Router			/question/update [post]
func (this QuestionController) UpdateQuestion() {
	var params question.QuestionUpdateRequest
	if err := this.BindJSON(&params); err != nil {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}

	// 判断是否存在
	questionObj, err := questionservice.GetById(params.Id)
	if err != nil || questionObj.Id <= 0 {
		myresq.Abort(this.Ctx, myresq.NOT_FOUND_ERROR, "题目未找到")
		return
	}

	if res := utils.CopyStructFields(params, questionObj); !res {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}
	if s, err := json.Marshal(params.Tags); err == nil && string(s) != "null" {
		questionObj.Tags = string(s)
	}
	if s, err := json.Marshal(params.JudgeCase); err == nil && string(s) != "null" {
		questionObj.JudgeCase = string(s)
	}
	if s, err := json.Marshal(params.JudgeConfig); err == nil && string(s) != "null" {
		questionObj.JudgeConfig = string(s)
	}
	// fmt.Printf("%+v\n", questionObj)

	// 参数校验
	questionservice.ValidQuestion(this.Ctx, questionObj, false)

	if err := questionservice.UpdateById(questionObj); err != nil {
		mylog.Log.Error("更新题目失败, err=", err.Error())
		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "更新失败")
		return
	}

	myresq.Success(this.Ctx, nil)
}

//	@Summary		根据 id 获取
//	@Description	根据 id 获取
//	@Tags			题目增删改查（管理员）
//	@Accept			application/x-www-form-urlencoded
//	@Produce		application/json
//	@Param			id	query		int							true	"id"
//	@Success		200	{object}	swagtype.QuestionResponse	"响应数据"
//	@Router			/question/get [get]
func (this QuestionController) GetQuestionById() {
	id, err := this.GetInt64("id")
	if err != nil || id <= 0 {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}
	questionObj, err := questionservice.GetById(id)
	if err != nil || questionObj.Id <= 0 {
		myresq.Abort(this.Ctx, myresq.NOT_FOUND_ERROR, "题目未找到")
		return
	}

	loginUserInterface := this.Ctx.Input.GetData("loginUser")
	loginUser, ok := loginUserInterface.(*entity.User)
	if !ok {
		myresq.Abort(this.Ctx, myresq.GET_CONTEXT_ERROR, "")
		return
	}

	// 不是本人或管理员，不能直接获取所有信息
	if !utils.CheckSame[int64]("检查当前用户与题目所属用户id是否一致", questionObj.UserId, loginUser.Id) && !userservice.IsAdmin(loginUser) {
		myresq.Abort(this.Ctx, myresq.NO_AUTH_ERROR, "")
		return
	}

	myresq.Success(this.Ctx, questionObj)
}

//	@Summary		根据 id 获取包装类（脱敏）
//	@Description	根据 id 获取包装类（脱敏）
//	@Tags			题目增删改查
//	@Accept			application/x-www-form-urlencoded
//	@Produce		application/json
//	@Param			id	query		int							true	"id"
//	@Success		200	{object}	swagtype.QuestionVOResponse	"响应数据"
//	@Router			/question/get/vo [get]
func (this QuestionController) GetQuestionVOById() {
	id, err := this.GetInt64("id")
	if err != nil || id <= 0 {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}
	questionObj, err := questionservice.GetById(id)
	if err != nil || questionObj.Id <= 0 {
		mylog.Log.Error("根据 id 获取题目包装类（脱敏）失败, err=", err.Error())
		myresq.Abort(this.Ctx, myresq.NOT_FOUND_ERROR, "题目未找到")
		return
	}

	// 脱敏
	respdata := questionservice.GetQuestionVO(this.Ctx, questionObj)
	myresq.Success(this.Ctx, respdata)
}

//	@Summary		分页获取列表
//	@Description	分页获取列表
//	@Tags			题目增删改查（管理员）
//	@Accept			application/json
//	@Produce		application/json
//	@Param			request	body		question.QuestionQueryRequest	true	"查询参数"
//	@Success		200		{object}	swagtype.ListQuestionResponse	"响应数据"
//	@Router			/question/list/page [post]
func (this QuestionController) ListQuestionByPage() {
	var params question.QuestionQueryRequest
	if err := this.BindJSON(&params); err != nil {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}

	// 限制爬虫
	if params.PageSize > 20 {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "超过分页大小限制")
		return
	}

	// 获取 QuerySeter 对象，直接使用 Model 结构体作为表名
	qs := mydb.O.QueryTable(new(entity.Question))

	// 构建查询条件
	qs = commonservice.GetQuerySeterByPage(qs, params.Current, params.PageSize)
	qs = questionservice.GetQuerySeter(qs, params)

	// 执行查询
	var questionPage []*entity.Question

	if _, err := qs.All(&questionPage); err != nil {
		mylog.Log.Errorf("ListQuestionByPage qs.All error: %v", err.Error())
		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "查询失败")
		return
	}

	num, err := questionservice.GetQuerySeter(mydb.O.QueryTable(new(entity.Question)), params).Count()
	if err != nil {
		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "查询失败")
		return
	}

	respdata := map[string]interface{}{
		"records": questionPage,
		"total":   num,
	}
	myresq.Success(this.Ctx, respdata)
}

//	@Summary		分页获取封装列表
//	@Description	分页获取封装列表
//	@Tags			题目增删改查
//	@Accept			application/json
//	@Produce		application/json
//	@Param			request	body		question.QuestionQueryRequest	true	"查询参数"
//	@Success		200		{object}	swagtype.ListQuestionVOResponse	"响应数据"
//	@Router			/question/list/page/vo [post]
func (this QuestionController) ListQuestionVOByPage() {
	var params question.QuestionQueryRequest
	if err := this.BindJSON(&params); err != nil {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}

	// 限制爬虫
	if params.PageSize > 20 {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "超过分页大小限制")
		return
	}

	// 获取 QuerySeter 对象，直接使用 Model 结构体作为表名
	qs := mydb.O.QueryTable(new(entity.Question))

	// 构建查询条件
	qs = commonservice.GetQuerySeterByPage(qs, params.Current, params.PageSize)
	qs = questionservice.GetQuerySeter(qs, params)

	// 执行查询
	var questionPage []*entity.Question

	if _, err := qs.All(&questionPage); err != nil {
		mylog.Log.Errorf("ListQuestionByPage qs.All error: %v", err.Error())
		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "查询失败")
		return
	}

	num, err := questionservice.GetQuerySeter(mydb.O.QueryTable(new(entity.Question)), params).Count()
	if err != nil {
		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "查询失败")
		return
	}

	respdata := map[string]interface{}{
		"records": questionservice.ListQuestionVO(this.Ctx, questionPage),
		"total":   num,
	}
	myresq.Success(this.Ctx, respdata)
}

//	@Summary		分页获取当前用户创建的资源列表
//	@Description	分页获取当前用户创建的资源列表
//	@Tags			题目增删改查
//	@Accept			application/json
//	@Produce		application/json
//	@Param			request	body		question.QuestionQueryRequest	true	"查询参数"
//	@Success		200		{object}	swagtype.ListQuestionVOResponse	"响应数据"
//	@Router			/question/my/list/page/vo [post]
func (this QuestionController) ListMyQuestionVOByPage() {
	var params question.QuestionQueryRequest
	if err := this.BindJSON(&params); err != nil {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}

	// 限制爬虫
	if params.PageSize > 20 {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "超过分页大小限制")
		return
	}

	// 获取 QuerySeter 对象，直接使用 Model 结构体作为表名
	qs := mydb.O.QueryTable(new(entity.Question))

	loginUserInterface := this.Ctx.Input.GetData("loginUser")
	loginUser, ok := loginUserInterface.(*entity.User)
	if !ok {
		myresq.Abort(this.Ctx, myresq.GET_CONTEXT_ERROR, "")
		return
	}

	// 构建查询条件
	qs = commonservice.GetQuerySeterByPage(qs, params.Current, params.PageSize)
	qs = questionservice.GetQuerySeter(qs, params)
	qs = qs.Filter("userId", loginUser.Id)

	// 执行查询
	var questionPage []*entity.Question

	if _, err := qs.All(&questionPage); err != nil {
		mylog.Log.Errorf("ListQuestionByPage qs.All error: %v", err.Error())
		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "查询失败")
		return
	}

	num, err := questionservice.GetQuerySeter(mydb.O.QueryTable(new(entity.Question)), params).Filter("userId", loginUser.Id).Count()
	if err != nil {
		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "查询失败")
		return
	}

	respdata := map[string]interface{}{
		"records": questionservice.ListQuestionVO(this.Ctx, questionPage),
		"total":   num,
	}
	myresq.Success(this.Ctx, respdata)
}

//	@Summary		提交题目
//	@Description	提交题目
//	@Tags			题目增删改查
//	@Accept			application/json
//	@Produce		application/json
//	@Param			request	body		questionsubmit.QuestionSubmitAddRequest	true	"提交参数"
//	@Success		200		{object}	object									"响应数据"
//	@Router			/question/question_submit/do [post]
func (this QuestionController) DoQuestionSubmit() {
	var params questionsubmit.QuestionSubmitAddRequest
	if err := this.BindJSON(&params); err != nil {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}
	// 登录才能提交
	loginUserInterface := this.Ctx.Input.GetData("loginUser")
	loginUser, ok := loginUserInterface.(*entity.User)
	if !ok {
		myresq.Abort(this.Ctx, myresq.GET_CONTEXT_ERROR, "")
		return
	}

	id := questionsubmitservice.DoQuestionSubmit(this.Ctx, params, loginUser)

	myresq.Success(this.Ctx, id)
}

//	@Summary		获取提交题目的封装
//	@Description	获取提交题目的封装（仅本人能看见自己提交的代码）
//	@Tags			题目增删改查
//	@Accept			application/x-www-form-urlencoded
//	@Produce		application/json
//	@Param			id	query		int									true	"id"
//	@Success		200	{object}	swagtype.QuestionSubmitVOResponse	"响应数据"
//	@Router			/question/question_submit/get/vo [get]
func (this QuestionController) GetQuestionSubmitVOById() {
	id, err := this.GetInt64("id")
	if err != nil || id <= 0 {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}

	// 登录才能查看
	loginUserInterface := this.Ctx.Input.GetData("loginUser")
	loginUser, ok := loginUserInterface.(*entity.User)
	if !ok {
		myresq.Abort(this.Ctx, myresq.GET_CONTEXT_ERROR, "")
		return
	}

	questionSubmitObj, err := questionsubmitservice.GetById(id)
	if err != nil {
		mylog.Log.Error("根据 id 获取已提交题目信息失败, err=", err.Error())
		myresq.Abort(this.Ctx, myresq.NOT_FOUND_ERROR, "题目未找到")
		return
	}

	// 脱敏
	respdata := questionsubmitservice.GetQuestionSubmitVO(this.Ctx, questionSubmitObj, loginUser)

	myresq.Success(this.Ctx, respdata)
}

//	@Summary		分页获取题目提交列表
//	@Description	分页获取题目提交列表（除了管理员外，普通用户只能看到非答案、提交代码等公开信息）
//	@Tags			题目增删改查
//	@Accept			application/json
//	@Produce		application/json
//	@Param			request	body		questionsubmit.QuestionSubmitQueryRequest	true	"查询参数"
//	@Success		200		{object}	swagtype.ListQuestionSubmitVOResponse		"响应数据"
//	@Router			/question/question_submit/list/page [post]
func (this QuestionController) ListQuestionSubmitByPage() {
	var params questionsubmit.QuestionSubmitQueryRequest
	if err := this.BindJSON(&params); err != nil {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}

	qs := mydb.O.QueryTable(new(entity.QuestionSubmit))
	qs = commonservice.GetQuerySeterByPage(qs, params.Current, params.PageSize)
	qs = questionsubmitservice.GetQuerySeter(qs, params)

	var questionSubmitPage []*entity.QuestionSubmit

	if _, err := qs.All(&questionSubmitPage); err != nil {
		mylog.Log.Errorf("ListQuestionSubmitByPage qs.All error: %v", err.Error())
		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "查询失败")
		return
	}

	num, err := questionsubmitservice.GetQuerySeter(mydb.O.QueryTable(new(entity.QuestionSubmit)), params).Count()
	if err != nil {
		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "查询失败")
		return
	}

	loginUserInterface := this.Ctx.Input.GetData("loginUser")
	loginUser, ok := loginUserInterface.(*entity.User)
	if !ok {
		myresq.Abort(this.Ctx, myresq.GET_CONTEXT_ERROR, "")
		return
	}

	questionSubmitVOPage := questionsubmitservice.ListQuestionSubmitVOPage(this.Ctx, questionSubmitPage, loginUser)

	respdata := map[string]interface{}{
		"records": questionSubmitVOPage,
		"total":   num,
	}
	myresq.Success(this.Ctx, respdata)
}
