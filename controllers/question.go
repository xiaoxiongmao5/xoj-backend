/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-26 22:18:34
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-27 20:53:40
 * @FilePath: /xoj-backend/controllers/question/question.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package controllers

import (
	"encoding/json"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/common"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/dto/question"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/entity"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/myresq"
	questionservice "github.com/xiaoxiongmao5/xoj/xoj-backend/service/questionService"
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

	var dbmodel entity.Question
	if res := utils.CopyStructFields(params, &dbmodel); !res {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}
	if s, err := json.Marshal(params.Tags); err == nil && string(s) != "null" {
		dbmodel.Tags = string(s)
	}
	if s, err := json.Marshal(params.Judgecase); err == nil && string(s) != "null" {
		dbmodel.Judgecase = string(s)
	}
	if s, err := json.Marshal(params.Judgeconfig); err == nil && string(s) != "null" {
		dbmodel.Judgeconfig = string(s)
	}

	// 参数校验
	questionservice.ValidQuestion(this.Ctx, &dbmodel, true)

	dbparams := question.QuestionAddRequest2DBParams(this.Ctx, &params)

	// loginUser := userservice.GetLoginUser(this.Ctx)
	// dbparams.Userid = loginUser.ID
	dbparams.Userid = 1

	res, err := questionservice.Save(dbparams)
	if err != nil {
		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "添加失败")
		return
	}

	respdata := make(map[string]int64)
	respdata["id"], _ = res.LastInsertId()
	myresq.Success(this.Ctx, respdata)
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
	if utils.CheckSame[int64]("检查当前用户与题目所属用户id是否一致", questionInfo.Userid, loginUser.ID) || !userservice.IsAdmin(loginUser) {
		myresq.Abort(this.Ctx, myresq.NO_AUTH_ERROR, "")
		return
	}

	if err = questionservice.RemoveById(params.ID); err != nil {
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

	var dbmodel entity.Question
	if res := utils.CopyStructFields(params, &dbmodel); !res {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}
	if s, err := json.Marshal(params.Tags); err == nil && string(s) != "null" {
		dbmodel.Tags = string(s)
	}
	if s, err := json.Marshal(params.Judgecase); err == nil && string(s) != "null" {
		dbmodel.Judgecase = string(s)
	}
	if s, err := json.Marshal(params.Judgeconfig); err == nil && string(s) != "null" {
		dbmodel.Judgeconfig = string(s)
	}

	// 参数校验
	questionservice.ValidQuestion(this.Ctx, &dbmodel, false)

	loginUser := userservice.GetLoginUser(this.Ctx)

	// 判断是否存在
	questionInfo, err := questionservice.GetById(params.ID)
	if err != nil || questionInfo.ID <= 0 {
		myresq.Abort(this.Ctx, myresq.NOT_FOUND_ERROR, "题目未找到")
		return
	}

	// 仅本人或管理员可编辑
	if utils.CheckSame[int64]("检查当前用户与题目所属用户id是否一致", questionInfo.Userid, loginUser.ID) || !userservice.IsAdmin(loginUser) {
		myresq.Abort(this.Ctx, myresq.NO_AUTH_ERROR, "")
		return
	}

	dbparams := question.QuestionEditRequest2DBParams(this.Ctx, &params)

	if err := questionservice.UpdateById(dbparams); err != nil {
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

	var dbmodel entity.Question
	if res := utils.CopyStructFields(params, &dbmodel); !res {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}
	if s, err := json.Marshal(params.Tags); err == nil && string(s) != "null" {
		dbmodel.Tags = string(s)
	}
	if s, err := json.Marshal(params.Judgecase); err == nil && string(s) != "null" {
		dbmodel.Judgecase = string(s)
	}
	if s, err := json.Marshal(params.Judgeconfig); err == nil && string(s) != "null" {
		dbmodel.Judgeconfig = string(s)
	}
	// fmt.Printf("%+v\n", dbmodel)

	// 参数校验
	questionservice.ValidQuestion(this.Ctx, &dbmodel, false)

	// 判断是否存在
	questionInfo, err := questionservice.GetById(params.ID)
	if err != nil || questionInfo.ID <= 0 {
		myresq.Abort(this.Ctx, myresq.NOT_FOUND_ERROR, "题目未找到")
		return
	}

	dbparams := question.QuestionUpdateRequest2DBParams(this.Ctx, &params)

	if err := questionservice.UpdateById(dbparams); err != nil {
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
	questionInfo, err := questionservice.GetById(id)
	if err != nil || questionInfo.ID <= 0 {
		myresq.Abort(this.Ctx, myresq.NOT_FOUND_ERROR, "题目未找到")
		return
	}
	// loginUser := userservice.GetLoginUser(this.Ctx)
	// // 不是本人或管理员，不能直接获取所有信息
	// if utils.CheckSame[int64]("检查当前用户与题目所属用户id是否一致", questionInfo.Userid, loginUser.ID) || !userservice.IsAdmin(loginUser) {
	// 	myresq.Abort(this.Ctx, myresq.NO_AUTH_ERROR, "")
	// 	return
	// }

	respdata := entity.DbConvertQuestion(questionInfo)
	myresq.Success(this.Ctx, respdata)
}

// 根据 id 获取（脱敏）
//
//	@Param			id	path/query		int								true	"题目id"
//	@router	/question/get/vo [get]
func (this QuestionController) GetQuestionVOById() {
	id, err := this.GetInt64("id")
	if err != nil || id <= 0 {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}
	questionInfo, err := questionservice.GetById(id)
	if err != nil || questionInfo.ID <= 0 {
		myresq.Abort(this.Ctx, myresq.NOT_FOUND_ERROR, "题目未找到")
		return
	}

	tmp := entity.DbConvertQuestion(questionInfo)
	// 脱敏
	respdata := questionservice.GetQuestionVO(tmp)
	myresq.Success(this.Ctx, respdata)
}

// 分页获取题目列表（仅管理员）
//
//	@router				/list/page [post]
//	@AuthCheck(mustRole	= UserConstant.ADMIN_ROLE)
func (this QuestionController) ListQuestionByPage() {

}

// 分页获取列表（封装类）
//
//	@router	/list/page/vo [post]
func (this QuestionController) ListQuestionVOByPage() {

}

// 分页获取当前用户创建的资源列表
//
//	@router	/my/list/page/vo [post]
func (this QuestionController) ListMyQuestionVOByPage() {

}

// 提交题目
//
//	@router	/question_submit/do [post]
func (this QuestionController) DoQuestionSubmit() {

}

// 分页获取题目提交列表（除了管理员外，普通用户只能看到非答案、提交代码等公开信息）
//
//	@router	/question_submit/list/page [post]
func (this QuestionController) ListQuestionSubmitByPage() {

}
