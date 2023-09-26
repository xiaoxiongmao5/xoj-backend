package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/models"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/mylog"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/myresq"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/service"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/store"
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

//	@Title			GetUserInfoByUserAccount
//	@Description	根据用户的 Cookie 获取用户信息
//	@Tags			用户相关
//	@Produce		application/json
//	@Success		200	{object}	object	"用户信息"
//	@router			/user/uinfo [get]
func (this UserController) GetUserInfoByUserAccount() {
	// 从请求中获取当前的 Token
	tokenCookie := this.Ctx.GetCookie("token")
	if tokenCookie == "" {
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}

	// 验证当前 Token
	token, err := jwt.Parse(tokenCookie, func(token *jwt.Token) (interface{}, error) {
		return []byte("SecretKey"), nil
	})
	if err != nil || !token.Valid {
		myresq.Abort(this.Ctx, myresq.NOT_LOGIN_ERROR, "")
		return
	}
	// 从 Token 中获取用户信息
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		myresq.Abort(this.Ctx, myresq.NOT_LOGIN_ERROR, "")
		return
	}

	// 重新生成 Token，并更新有效期
	userAccount := claims["user_account"].(string)
	userRole := claims["user_role"].(string)
	newToken, err := utils.GenerateToken(userAccount, userRole)
	if err != nil {
		myresq.Abort(this.Ctx, myresq.GENERATE_TOKEN_FAILED, err.Error())
		return
	}

	// 删除旧的 token
	delete(store.TokenMemoryStore, tokenCookie)

	// 更新内存中的 token 数据
	store.TokenMemoryStore[newToken] = true

	// 将新的 token 返回给前端
	domain, _ := utils.GetDomainFromReferer(this.Ctx.Request.Referer())
	this.Ctx.SetCookie("token", newToken, 3600, "/", domain, false, true)

	useraccount := claims["user_account"]
	// if useraccount == "" {
	// 	mylog.Log.Error("从上下文获取user_account失败")
	// 	myresq.Abort(this.Ctx, myresq.GET_CONTEXT_ERROR, "")
	// 	return
	// }

	userInfo, err := service.GetUserInfoByUserAccount(useraccount.(string))
	if err != nil {
		mylog.Log.Errorf("q.GetUserInfoByUserAccount err=%v", err.Error())
		myresq.Abort(this.Ctx, myresq.USER_NOT_EXIST, "")
		return
	}

	myresq.Success(this.Ctx, models.ConvertToNormalUser(userInfo))
}

//	@Title			Register
//	@Description	用户注册
//	@Tags			用户相关
//	@Accept			application/json
//	@Produce		application/json
//	@Param			request	body		models.UserRegisterParams	true	"注册信息"
//	@Success		200		{object}	object
//	@router			/user/register [post]
func (this UserController) Register() {
	var params models.UserRegisterParams
	if err := this.BindJSON(&params); err != nil {
		mylog.Log.Errorf("UserRegisterParams err=%v", err.Error())
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}

	// 注册用户
	if _, err := service.CreateUser(&params); err != nil {
		mylog.Log.Errorf("service.CreateUser err=%v", err.Error())
		myresq.Abort(this.Ctx, myresq.CREATE_USER_FAILED, err.Error())
		return
	}

	myresq.Success(this.Ctx, nil)
}

//	@Title			Login
//	@Summary		用户登录
//	@Description	用户登录
//	@Tags			用户相关
//	@Accept			application/x-www-form-urlencoded
//	@Produce		application/json
//	@Param			request	formData	models.UserLoginParams	true	"账号密码"
//	@Success		200		{object}	object
//	@router			/user/login [post]
func (this UserController) Login() {
	// 检查用户是否已经登录
	// tokenCookie := this.Ctx.GetCookie("token")
	// if tokenCookie != "" {
	// 已经登录，直接返回登录成功
	// ghandle.HandlerSuccess(c, "Already logged in", nil)
	// return
	// }
	// service.DeleteToken(ctx)

	var params models.UserLoginParams
	if err := this.BindForm(&params); err != nil {
		mylog.Log.Errorf("UserLoginParams err=%v", err.Error())
		myresq.Abort(this.Ctx, myresq.PARAMS_ERROR, "")
		return
	}

	// 用户是否存在（获取用户信息）
	userInfo, err := service.GetUserInfoByUserAccount(params.Useraccount)
	if err != nil {
		mylog.Log.Errorf("q.GetUserInfoByUserAccount err=%v", err.Error())
		myresq.Abort(this.Ctx, myresq.USER_NOT_EXIST, "")
		return
	}
	mylog.Log.Infof("拿到用户信息了%v", userInfo)

	// 验证密码是否正常
	if err := utils.CheckHashPasswordByBcrypt(userInfo.Userpassword, params.Userpassword); err != nil {
		mylog.Log.Errorf("CheckHashPasswordByBcrypt err=%v", err.Error())
		myresq.Abort(this.Ctx, myresq.USER_PASSWORD_ERROR, "账号不存在或者密码验证错误")
		return
	}

	// 生成token
	token, err := utils.GenerateToken(params.Useraccount, userInfo.Userrole)
	if err != nil {
		mylog.Log.Errorf("utils.GenerateToken err=%v", err.Error())
		myresq.Abort(this.Ctx, myresq.GENERATE_TOKEN_FAILED, err.Error())
		return
	}

	// 存储token
	store.TokenMemoryStore[token] = true

	// 返回token到前端
	referer := this.Ctx.Request.Referer()
	mylog.Log.Info("refer=", referer) //http://localhost:8080/
	domain, _ := utils.GetDomainFromReferer(referer)
	mylog.Log.Info("domain=", domain) //localhost
	this.Ctx.SetCookie("token", token, 3600, "/", domain, false, true)

	myresq.Success(this.Ctx, models.ConvertToNormalUser(userInfo))
}

//	@Title			logout
//	@Description	Logs out current logged in user session
//	@Success		200	{string}	logout	success
//	@router			/logout [get]
func (this UserController) Logout() {
	this.Data["json"] = "logout success"
	this.ServeJSON()
}
