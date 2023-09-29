/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-26 22:18:34
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-29 23:08:01
 * @FilePath: /xoj-backend/controllers/question/question.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
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

// 创建
//
//	@Param	request	body	question.QuestionAddRequest	true	"注册信息"
//	@router	/question/add [post]
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
	if s, err := json.Marshal(params.Judgecase); err == nil && string(s) != "null" {
		questionObj.Judgecase = string(s)
	}
	if s, err := json.Marshal(params.Judgeconfig); err == nil && string(s) != "null" {
		questionObj.Judgeconfig = string(s)
	}

	// 参数校验
	questionservice.ValidQuestion(this.Ctx, &questionObj, true)

	loginUser := userservice.GetLoginUser(this.Ctx)
	questionObj.Userid = loginUser.ID

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

// 删除
//
//	@Param	request	body	common.IdRequest	true	"题目id"
//	@router	/question/delete [post]
func (this QuestionController) DeleteQuestion() {
	var params common.IdRequest
	if err := this.BindJSON(&params); err != nil {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}

	loginUser := userservice.GetLoginUser(this.Ctx)

	// 判断是否存在
	questionInfo, err := questionservice.GetById(params.ID)
	if err != nil || questionInfo.ID <= 0 {
		myresq.Abort(this.Ctx, myresq.NOT_FOUND_ERROR, "题目未找到")
		return
	}
	// 仅本人或管理员可删除
	if !utils.CheckSame[int64]("检查当前用户与题目所属用户id是否一致", questionInfo.Userid, loginUser.ID) && !userservice.IsAdmin(loginUser) {
		myresq.Abort(this.Ctx, myresq.NO_AUTH_ERROR, "")
		return
	}

	if err = questionservice.RemoveById(params.ID); err != nil {
		mylog.Log.Error("删除题目失败, err=", err.Error())
		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "删除失败")
	}

	myresq.Success(this.Ctx, nil)
}

// 编辑（用户）
//
//	@router	/question/edit [post]
func (this QuestionController) EditQuestion() {
	var params question.QuestionEditRequest
	if err := this.BindJSON(&params); err != nil {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}

	// 判断是否存在
	questionObj, err := questionservice.GetById(params.ID)
	if err != nil || questionObj.ID <= 0 {
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
	if s, err := json.Marshal(params.Judgecase); err == nil && string(s) != "null" {
		questionObj.Judgecase = string(s)
	}
	if s, err := json.Marshal(params.Judgeconfig); err == nil && string(s) != "null" {
		questionObj.Judgeconfig = string(s)
	}

	// 参数校验
	questionservice.ValidQuestion(this.Ctx, questionObj, false)

	loginUser := userservice.GetLoginUser(this.Ctx)

	// 仅本人或管理员可编辑
	if !utils.CheckSame[int64]("检查当前用户与题目所属用户id是否一致", questionObj.Userid, loginUser.ID) && !userservice.IsAdmin(loginUser) {
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

// 更新（仅管理员）
//
//	@router				/question/update [post]
//	@AuthCheck(mustRole	= UserConstant.ADMIN_ROLE)
func (this QuestionController) UpdateQuestion() {
	var params question.QuestionUpdateRequest
	if err := this.BindJSON(&params); err != nil {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}

	// 判断是否存在
	questionObj, err := questionservice.GetById(params.ID)
	if err != nil || questionObj.ID <= 0 {
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
	if s, err := json.Marshal(params.Judgecase); err == nil && string(s) != "null" {
		questionObj.Judgecase = string(s)
	}
	if s, err := json.Marshal(params.Judgeconfig); err == nil && string(s) != "null" {
		questionObj.Judgeconfig = string(s)
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

// 根据 id 获取
//
//	@Param			id	path/query		int								true	"题目id"
//	@router	/question/get [get]
func (this QuestionController) GetQuestionById() {
	id, err := this.GetInt64("id")
	if err != nil || id <= 0 {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}
	questionObj, err := questionservice.GetById(id)
	if err != nil || questionObj.ID <= 0 {
		myresq.Abort(this.Ctx, myresq.NOT_FOUND_ERROR, "题目未找到")
		return
	}

	loginUser := userservice.GetLoginUser(this.Ctx)

	// 不是本人或管理员，不能直接获取所有信息
	if !utils.CheckSame[int64]("检查当前用户与题目所属用户id是否一致", questionObj.Userid, loginUser.ID) && !userservice.IsAdmin(loginUser) {
		myresq.Abort(this.Ctx, myresq.NO_AUTH_ERROR, "")
		return
	}

	myresq.Success(this.Ctx, questionObj)
}

// 根据 id 获取包装类（脱敏）
//
//	@Param			id	path/query		int								true	"题目id"
//	@router	/question/get/vo [get]
func (this QuestionController) GetQuestionVOById() {
	id, err := this.GetInt64("id")
	if err != nil || id <= 0 {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}
	questionObj, err := questionservice.GetById(id)
	if err != nil || questionObj.ID <= 0 {
		mylog.Log.Error("根据 id 获取题目包装类（脱敏）失败, err=", err.Error())
		myresq.Abort(this.Ctx, myresq.NOT_FOUND_ERROR, "题目未找到")
		return
	}

	// 脱敏
	respdata := questionservice.GetQuestionVO(this.Ctx, questionObj)
	myresq.Success(this.Ctx, respdata)
}

// 分页获取题目列表（仅管理员）
//
//	@router				/question/list/page [post]
//	@AuthCheck(mustRole	= UserConstant.ADMIN_ROLE)
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

	num, err := qs.All(&questionPage)
	if err != nil {
		mylog.Log.Errorf("ListQuestionByPage qs.All error: %v", err.Error())
		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "查询失败")
		return
	}

	respdata := map[string]interface{}{
		"data":  questionPage,
		"total": num,
	}
	myresq.Success(this.Ctx, respdata)
}

// 分页获取题目列表（封装类）
//
//	@router	/question/list/page/vo [post]
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

	num, err := qs.All(&questionPage)
	if err != nil {
		mylog.Log.Errorf("ListQuestionByPage qs.All error: %v", err.Error())
		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "查询失败")
		return
	}

	respdata := map[string]interface{}{
		"data":  questionservice.ListQuestionVO(this.Ctx, questionPage),
		"total": num,
	}
	myresq.Success(this.Ctx, respdata)
}

// 分页获取当前用户创建的资源列表
//
//	@router	/question/my/list/page/vo [post]
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

	loginUser := userservice.GetLoginUser(this.Ctx)

	// 构建查询条件
	qs = commonservice.GetQuerySeterByPage(qs, params.Current, params.PageSize)
	qs = questionservice.GetQuerySeter(qs, params)
	qs = qs.Filter("userid", loginUser.ID)

	// 执行查询
	var questionPage []*entity.Question

	num, err := qs.All(&questionPage)
	if err != nil {
		mylog.Log.Errorf("ListQuestionByPage qs.All error: %v", err.Error())
		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "查询失败")
		return
	}

	respdata := map[string]interface{}{
		"data":  questionservice.ListQuestionVO(this.Ctx, questionPage),
		"total": num,
	}
	myresq.Success(this.Ctx, respdata)
}

// 提交题目
//
//	@router	/question_submit/do [post]
func (this QuestionController) DoQuestionSubmit() {
	var params questionsubmit.QuestionSubmitAddRequest
	if err := this.BindJSON(&params); err != nil {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}
	// 登录才能提交
	loginUser := userservice.GetLoginUser(this.Ctx)

	id := questionsubmitservice.DoQuestionSubmit(this.Ctx, params, loginUser)

	myresq.Success(this.Ctx, id)
}

// 分页获取题目提交列表（除了管理员外，普通用户只能看到非答案、提交代码等公开信息）
//
//	@router	/question_submit/list/page [post]
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

	num, err := qs.All(&questionSubmitPage)
	if err != nil {
		mylog.Log.Errorf("ListQuestionSubmitByPage qs.All error: %v", err.Error())
		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "查询失败")
		return
	}

	loginUser := userservice.GetLoginUser(this.Ctx)

	questionSubmitVOPage := questionsubmitservice.ListQuestionSubmitVOPage(this.Ctx, questionSubmitPage, loginUser)

	respdata := map[string]interface{}{
		"data":  questionSubmitVOPage,
		"total": num,
	}
	myresq.Success(this.Ctx, respdata)
}
