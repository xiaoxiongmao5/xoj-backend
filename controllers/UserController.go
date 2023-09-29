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

// 用户注册
// /register
func (this UserController) UserRegister() {
	var params user.UserRegisterRequest
	if err := this.BindJSON(&params); err != nil {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}

	id := userservice.UserRegister(this.Ctx, params.UserAccount, params.UserPassword, params.CheckUserpassword)

	myresq.Success(this.Ctx, id)
}

// 用户登录
// /login
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
// /login/wx_open
func (this UserController) UserLoginByWxOpen() {

}

// 用户注销
// /logout
func (this UserController) UserLogout() {
	ok := userservice.UserLogout(this.Ctx)
	if !ok {
		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "退出失败")
		return
	}
	myresq.Success(this.Ctx, nil)
}

// 获取当前登录用户
// /get/login
func (this UserController) GetLoginUser() {
	loginUser := userservice.GetLoginUser(this.Ctx)

	myresq.Success(this.Ctx, userservice.GetLoginUserVO(loginUser))
}

// 增删改查（管理员）
// /add
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

// 删除用户（管理员）
// /delete
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

// 更新用户（管理员）
// /update
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

// 根据 id 获取用户（仅管理员）
//
//	@Param			id	path/query		int								true	"用户id"
//
// /get [get]
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

// 根据 id 获取包装类（脱敏）
//
//	@Param			id	path/query		int								true	"用户id"
//
// /get/vo	[get]
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

// 分页获取用户列表（仅管理员）
// /list/page  [post]
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

	num, err := qs.All(&userPage)
	if err != nil {
		mylog.Log.Errorf("ListUserByPage qs.All error: %v", err.Error())
		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "查询失败")
		return
	}

	respdata := map[string]interface{}{
		"data":  userPage,
		"total": num,
	}
	myresq.Success(this.Ctx, respdata)

}

// 分页获取用户封装列表
// /list/page/vo [post]
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

	num, err := qs.All(&userPage)
	if err != nil {
		mylog.Log.Errorf("ListUserByPage qs.All error: %v", err.Error())
		myresq.Abort(this.Ctx, myresq.OPERATION_ERROR, "查询失败")
		return
	}

	respdata := map[string]interface{}{
		"data":  userservice.ListUserVO(userPage),
		"total": num,
	}
	myresq.Success(this.Ctx, respdata)

}

// 更新个人信息
// /update/my [post]
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

// @Title			GetUserInfoByUserAccount
// @Description	根据用户的 Cookie 获取用户信息
// @Tags			用户相关
// @Produce		application/json
// @Success		200	{object}	object	"用户信息"
// @router			/user/uinfo [get]
// func (this UserController) GetUserInfoByUserAccount() {
// 	loginUser := service.GetLoginUser(this.Ctx)
// 	myresq.Success(this.Ctx, models.ConvertToNormalUser(loginUser))
// }

// @Title			Register
// @Description	用户注册
// @Tags			用户相关
// @Accept			application/json
// @Produce		application/json
// @Param			request	body		user.UserRegisterRequest	true	"注册信息"
// @Success		200		{object}	object
// @router			/user/register [post]
// func (this UserController) Register() {
// 	var params user.UserRegisterRequest
// 	if err := this.BindJSON(&params); err != nil {
// 		mylog.Log.Errorf("UserRegisterParams err=%v", err.Error())
// 		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
// 		return
// 	}

// 	// 注册用户
// 	if _, err := service.CreateUser(&params); err != nil {
// 		mylog.Log.Errorf("service.CreateUser err=%v", err.Error())
// 		myresq.Abort(this.Ctx, myresq.CREATE_USER_FAILED, err.Error())
// 		return
// 	}

// 	myresq.Success(this.Ctx, nil)
// }

// @Title			Login
// @Summary		用户登录
// @Description	用户登录
// @Tags			用户相关
// @Accept			application/x-www-form-urlencoded
// @Produce		application/json
// @Param			request	formData	user.UserLoginRequest	true	"账号密码"
// @Success		200		{object}	object
// @router			/user/login [post]
// func (this UserController) Login() {
// 	// 检查用户是否已经登录
// 	// tokenCookie := this.Ctx.GetCookie("token")
// 	// if tokenCookie != "" {
// 	// 已经登录，直接返回登录成功
// 	// ghandle.HandlerSuccess(c, "Already logged in", nil)
// 	// return
// 	// }
// 	// service.DeleteToken(ctx)

// 	var params user.UserLoginRequest
// 	if err := this.BindForm(&params); err != nil {
// 		mylog.Log.Errorf("UserLoginParams err=%v", err.Error())
// 		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
// 		return
// 	}

// 	// 用户是否存在（获取用户信息）
// 	userInfo, err := service.GetUserInfoByUserAccount(params.UserAccount)
// 	if err != nil {
// 		mylog.Log.Errorf("q.GetUserInfoByUserAccount err=%v", err.Error())
// 		myresq.Abort(this.Ctx, myresq.USER_NOT_EXIST, "")
// 		return
// 	}
// 	mylog.Log.Infof("拿到用户信息了%v", userInfo)

// 	// 验证密码是否正常
// 	if err := utils.CheckHashPasswordByBcrypt(userInfo.UserPassword, params.UserPassword); err != nil {
// 		mylog.Log.Errorf("CheckHashPasswordByBcrypt err=%v", err.Error())
// 		myresq.Abort(this.Ctx, myresq.USER_PASSWORD_ERROR, "账号不存在或者密码验证错误")
// 		return
// 	}

// 	// 生成token
// 	token, err := utils.GenerateToken(params.UserAccount, userInfo.UserRole)
// 	if err != nil {
// 		mylog.Log.Errorf("utils.GenerateToken err=%v", err.Error())
// 		myresq.Abort(this.Ctx, myresq.GENERATE_TOKEN_FAILED, err.Error())
// 		return
// 	}

// 	// 存储token
// 	store.TokenMemoryStore[token] = true

// 	// 返回token到前端
// 	referer := this.Ctx.Request.Referer()
// 	mylog.Log.Info("refer=", referer) //http://localhost:8080/
// 	domain, _ := utils.GetDomainFromReferer(referer)
// 	mylog.Log.Info("domain=", domain) //localhost
// 	this.Ctx.SetCookie("token", token, 3600, "/", domain, false, true)

// 	myresq.Success(this.Ctx, models.ConvertToNormalUser(userInfo))
// }

// @Title			logout
// @Description	Logs out current logged in user session
// @Success		200	{string}	logout	success
// @router			/logout [get]
func (this UserController) Logout() {
	this.Data["json"] = "logout success"
	this.ServeJSON()
}
