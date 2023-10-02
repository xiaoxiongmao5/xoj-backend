package controllers

import (
	"fmt"

	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/common"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/dto/user"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/model/entity"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mydb"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mylog"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/myresq"
	commonservice "github.com/xiaoxiongmao5/xoj/xoj-backend/service/commonService"
	userservice "github.com/xiaoxiongmao5/xoj/xoj-backend/service/userService"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/utils"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// 在每一个请求执行之前调用一遍
func (this UserController) Prepare() {
	// this.EnableXSRF = false
	// this.XSRFExpire = 7200
}

// 在每一个请求执行之后调用一遍
func (this UserController) Finish() {
}

//	@Summary		用户注册
//	@Description	用户注册
//	@Tags			用户相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			request	body		user.UserRegisterRequest	true	"注册信息"
//	@Success		200		{object}	object						"响应数据"
//	@Router			/user/register [post]
func (this UserController) UserRegister() {
	var params user.UserRegisterRequest
	if err := this.BindJSON(&params); err != nil {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}

	id := userservice.UserRegister(this.Ctx, params.UserAccount, params.UserPassword, params.CheckUserpassword)

	myresq.Success(this.Ctx, id)
}

//	@Summary		用户登录
//	@Description	用户登录
//	@Tags			用户相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			request	body		user.UserLoginRequest			true	"账号密码"
//	@Success		200		{object}	swagtype.LoginUserVOResponse	"响应数据"
//	@Router			/user/login [post]
func (this UserController) UserLogin() {
	var params user.UserLoginRequest
	if err := this.BindJSON(&params); err != nil {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}
	loginUserVO := userservice.UserLogin(this.Ctx, params.UserAccount, params.UserPassword)

	myresq.Success(this.Ctx, loginUserVO)
}

// 用户登录（微信开放平台）
// /user/login/wx_open
func (this UserController) UserLoginByWxOpen() {}

//	@Summary		用户退出
//	@Description	用户退出
//	@Tags			用户相关
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200	{object}	object	"响应数据"
//	@Router			/user/logout [post]
func (this UserController) UserLogout() {
	ok := userservice.UserLogout(this.Ctx)
	if !ok {
		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "退出失败")
		return
	}
	myresq.Success(this.Ctx, nil)
}

//	@Summary		获取当前登录用户
//	@Description	获取当前登录用户
//	@Tags			用户相关
//	@Produce		application/json
//	@Success		200	{object}	swagtype.LoginUserVOResponse	"响应数据"
//	@Router			/user/get/login [get]
func (this UserController) GetLoginUser() {
	loginUser := userservice.GetLoginUser(this.Ctx)

	myresq.Success(this.Ctx, userservice.GetLoginUserVO(loginUser))
}

//	@Summary		添加
//	@Description	添加
//	@Tags			用户增删改查（管理员）
//	@Accept			application/json
//	@Produce		application/json
//	@Param			request	body		user.UserAddRequest	true	"添加参数"
//	@Success		200		{object}	object				"响应数据"
//	@Router			/user/add [post]
func (this UserController) AddUser() {
	var params user.UserAddRequest
	if err := this.BindJSON(&params); err != nil {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}

	userObj := entity.User{}
	if res := utils.CopyStructFields(params, &userObj); !res {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}
	id, err := userservice.Save(&userObj)
	if err != nil {
		mylog.Log.Error("添加用户失败, err=", err.Error())
		errmsg := fmt.Sprintf("添加用户失败, err=%v", err.Error())
		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, errmsg)
		return
	}

	myresq.Success(this.Ctx, id)
}

//	@Summary		删除
//	@Description	删除
//	@Tags			用户增删改查（管理员）
//	@Accept			application/json
//	@Produce		application/json
//	@Param			request	body		common.IdRequest	true	"删除id"
//	@Success		200		{object}	object				"响应数据"
//	@Router			/user/delete [post]
func (this UserController) DeleteUser() {
	var params common.IdRequest
	if err := this.BindJSON(&params); err != nil {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}

	if params.ID <= 0 {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}

	err := userservice.RemoveById(params.ID)
	if err != nil {
		mylog.Log.Error("删除用户失败, err=", err.Error())
		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "删除失败")
	}

	myresq.Success(this.Ctx, nil)
}

//	@Summary		更新
//	@Description	更新
//	@Tags			用户增删改查（管理员）
//	@Accept			application/json
//	@Produce		application/json
//	@Param			request	body		user.UserUpdateRequest	true	"更新参数"
//	@Success		200		{object}	object					"响应数据"
//	@Router			/user/update [post]
func (this UserController) UpdateUser() {
	var params user.UserUpdateRequest
	if err := this.BindJSON(&params); err != nil {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}

	if params.ID <= 0 {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}

	userObj, err := userservice.GetById(params.ID)
	if err != nil {
		myresq.Abort(this.Ctx, myresq.USER_NOT_EXIST, "")
		return
	}
	utils.CopyStructFields(params, userObj)

	if err := userservice.UpdateById(userObj); err != nil {
		mylog.Log.Error("更新用户失败, err=", err.Error())
		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "更新失败")
	}

	myresq.Success(this.Ctx, nil)
}

//	@Summary		根据 id 获取
//	@Description	根据 id 获取
//	@Tags			用户增删改查（管理员）
//	@Accept			application/x-www-form-urlencoded
//	@Produce		application/json
//	@Param			id	query		int						true	"id"
//	@Success		200	{object}	swagtype.UserResponse	"响应数据"
//	@Router			/user/get [get]
func (this UserController) GetUserById() {
	id, err := this.GetInt64("id")
	if err != nil || id <= 0 {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}
	userObj, err := userservice.GetById(id)
	if err != nil {
		mylog.Log.Error("根据 id 获取用户（仅管理员）失败, err=", err.Error())
		myresq.Abort(this.Ctx, myresq.NOT_FOUND_ERROR, "")
	}

	myresq.Success(this.Ctx, userObj)
}

//	@Summary		根据 id 获取包装类（脱敏）
//	@Description	根据 id 获取包装类（脱敏）
//	@Tags			用户增删改查
//	@Accept			application/x-www-form-urlencoded
//	@Produce		application/json
//	@Param			id	query		int						true	"id"
//	@Success		200	{object}	swagtype.UserVOResponse	"响应数据"
//	@Router			/user/get/vo [get]
func (this UserController) GetUserVOById() {
	id, err := this.GetInt64("id")
	if err != nil || id <= 0 {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}
	userObj, err := userservice.GetById(id)
	if err != nil {
		mylog.Log.Error("根据 id 获取包装类（脱敏）失败, err=", err.Error())
		myresq.Abort(this.Ctx, myresq.NOT_FOUND_ERROR, "")
	}

	myresq.Success(this.Ctx, userservice.GetUserVO(userObj))
}

//	@Summary		分页获取列表
//	@Description	分页获取列表
//	@Tags			用户增删改查（管理员）
//	@Accept			application/json
//	@Produce		application/json
//	@Param			request	body		user.UserQueryRequest		true	"查询参数"
//	@Success		200		{object}	swagtype.ListUserResponse	"响应数据"
//	@Router			/user/list/page [post]
func (this UserController) ListUserByPage() {
	var params user.UserQueryRequest
	if err := this.BindJSON(&params); err != nil {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}

	qs := mydb.O.QueryTable(new(entity.User))

	qs = commonservice.GetQuerySeterByPage(qs, params.Current, params.PageSize)
	qs = userservice.GetQuerySeter(qs, params)

	var userPage []*entity.User

	if _, err := qs.All(&userPage); err != nil {
		mylog.Log.Errorf("ListUserByPage qs.All error: %v", err.Error())
		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "查询失败")
		return
	}

	num, err := userservice.GetQuerySeter(mydb.O.QueryTable(new(entity.User)), params).Count()
	if err != nil {
		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "查询失败")
		return
	}

	respdata := map[string]interface{}{
		"records": userPage,
		"total":   num,
	}
	myresq.Success(this.Ctx, respdata)

}

//	@Summary		分页获取封装列表
//	@Description	分页获取封装列表
//	@Tags			用户增删改查
//	@Accept			application/json
//	@Produce		application/json
//	@Param			request	body		user.UserQueryRequest		true	"查询参数"
//	@Success		200		{object}	swagtype.ListUserVOResponse	"响应数据"
//	@Router			/user/list/page/vo [post]
func (this UserController) ListUserVOByPage() {
	var params user.UserQueryRequest
	if err := this.BindJSON(&params); err != nil {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}

	// 限制爬虫
	if params.PageSize > 20 {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "超过分页大小限制")
		return
	}

	qs := mydb.O.QueryTable(new(entity.User))

	qs = commonservice.GetQuerySeterByPage(qs, params.Current, params.PageSize)
	qs = userservice.GetQuerySeter(qs, params)

	var userPage []*entity.User

	if _, err := qs.All(&userPage); err != nil {
		mylog.Log.Errorf("ListUserByPage qs.All error: %v", err.Error())
		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "查询失败")
		return
	}

	num, err := userservice.GetQuerySeter(mydb.O.QueryTable(new(entity.User)), params).Count()
	if err != nil {
		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "查询失败")
		return
	}

	respdata := map[string]interface{}{
		"records": userservice.ListUserVO(userPage),
		"total":   num,
	}
	myresq.Success(this.Ctx, respdata)

}

//	@Summary		更新个人信息
//	@Description	更新个人信息
//	@Tags			用户增删改查
//	@Accept			application/json
//	@Produce		application/json
//	@Param			request	body		user.UserUpdateMyRequest	true	"查询参数"
//	@Success		200		{object}	object						"响应数据"
//	@Router			/user/update/my [post]
func (this UserController) UpdateMyUser() {
	var params user.UserUpdateMyRequest
	if err := this.BindJSON(&params); err != nil {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}

	loginUser := userservice.GetLoginUser(this.Ctx)

	utils.CopyStructFields(params, loginUser)

	if err := userservice.UpdateById(loginUser); err != nil {
		mylog.Log.Error("更新个人信息失败, err=", err.Error())
		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "更新个人信息失败")
	}

	myresq.Success(this.Ctx, nil)
}
