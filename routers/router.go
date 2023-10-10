/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-26 10:35:03
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-10-10 22:51:47
 * @FilePath: /xoj-backend/routers/router.go
 */
//	@APIVersion			1.0.0
//	@Title				beego Test API
//	@Description		beego has a very cool tools to autogenerate documents for your API
//	@Contact			astaxie@gmail.com
//	@TermsOfServiceUrl	http://beego.me/
//	@License			Apache 2.0
//	@LicenseUrl			http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/controllers"
	"github.com/xiaoxiongmao5/xoj/xoj-backend/middleware"
)

func init() {
	// beego.AutoRouter(&controllers.UserController{})

	// beego.CtrlPost("/user/register", controllers.UserController.UserRegister)
	// beego.CtrlPost("/user/login", controllers.UserController.UserLogin)
	// beego.CtrlPost("/user/logout", controllers.UserController.UserLogout)
	// beego.CtrlGet("/user/get/login", controllers.UserController.GetLoginUser)
	// beego.CtrlGet("/user/get/vo", controllers.UserController.GetUserVOById)
	// beego.CtrlPost("/user/list/page/vo", controllers.UserController.ListUserVOByPage)
	// beego.CtrlGet("/question/get/vo", controllers.QuestionController.GetQuestionVOById)
	// beego.CtrlPost("/question/list/page/vo", controllers.QuestionController.ListQuestionVOByPage)

	beego.InsertFilter("*", beego.BeforeRouter, middleware.LogMiddleware)
	beego.InsertFilter("*", beego.AfterExec, middleware.LogMiddlewareAfter)
	beego.InsertFilter("*", beego.BeforeRouter, middleware.CORSMiddleware())
	beego.InsertFilter("*", beego.BeforeRouter, middleware.IPRateLimiterMiddleware)

	// ———— 需要登录
	// 用户相关
	beego.CtrlPost("/user/update/my", controllers.UserController.UpdateMyUser)
	beego.InsertFilter("/user/update/my", beego.BeforeRouter, middleware.AuthMiddleware)

	// 题目相关
	beego.CtrlPost("/question/question_submit/list/page", controllers.QuestionController.ListQuestionSubmitByPage)
	beego.InsertFilter("/question/question_submit/list/page", beego.BeforeRouter, middleware.AuthMiddleware)
	beego.CtrlPost("/question/question_submit/do", controllers.QuestionController.DoQuestionSubmit)
	beego.InsertFilter("/question/question_submit/do", beego.BeforeRouter, middleware.AuthMiddleware)
	beego.CtrlPost("/question/my/list/page/vo", controllers.QuestionController.ListMyQuestionVOByPage)
	beego.InsertFilter("/question/my/list/page/vo", beego.BeforeRouter, middleware.AuthMiddleware)
	beego.CtrlGet("/question/question_submit/get/vo", controllers.QuestionController.GetQuestionSubmitVOById)
	beego.InsertFilter("/question/question_submit/get/vo", beego.BeforeRouter, middleware.AuthMiddleware)

	// ———— 需要管理员
	// 用户相关
	beego.CtrlPost("/user/add", controllers.UserController.AddUser)
	beego.InsertFilter("/user/add", beego.BeforeRouter, middleware.AuthMiddleware)
	beego.InsertFilter("/user/add", beego.BeforeRouter, middleware.AdminMiddleware)
	beego.CtrlPost("/user/delete", controllers.UserController.DeleteUser)
	beego.InsertFilter("/user/delete", beego.BeforeRouter, middleware.AuthMiddleware)
	beego.InsertFilter("/user/delete", beego.BeforeRouter, middleware.AdminMiddleware)
	beego.CtrlPost("/user/update", controllers.UserController.UpdateUser)
	beego.InsertFilter("/user/update", beego.BeforeRouter, middleware.AuthMiddleware)
	beego.InsertFilter("/user/update", beego.BeforeRouter, middleware.AdminMiddleware)
	beego.CtrlGet("/user/get", controllers.UserController.GetUserById)
	beego.InsertFilter("/user/get", beego.BeforeRouter, middleware.AuthMiddleware)
	beego.InsertFilter("/user/get", beego.BeforeRouter, middleware.AdminMiddleware)
	beego.CtrlPost("/user/list/page", controllers.UserController.ListUserByPage)
	beego.InsertFilter("/user/list/page", beego.BeforeRouter, middleware.AuthMiddleware)
	beego.InsertFilter("/user/list/page", beego.BeforeRouter, middleware.AdminMiddleware)

	// 题目相关
	beego.CtrlPost("/question/add", controllers.QuestionController.AddQuestion)
	beego.InsertFilter("/question/add", beego.BeforeRouter, middleware.AuthMiddleware)
	beego.InsertFilter("/question/add", beego.BeforeRouter, middleware.AdminMiddleware)
	beego.CtrlPost("/question/delete", controllers.QuestionController.DeleteQuestion)
	beego.InsertFilter("/question/delete", beego.BeforeRouter, middleware.AuthMiddleware)
	beego.InsertFilter("/question/delete", beego.BeforeRouter, middleware.AdminMiddleware)
	beego.CtrlPost("/question/edit", controllers.QuestionController.EditQuestion)
	beego.InsertFilter("/question/edit", beego.BeforeRouter, middleware.AuthMiddleware)
	beego.InsertFilter("/question/edit", beego.BeforeRouter, middleware.AdminMiddleware)
	beego.CtrlGet("/question/get", controllers.QuestionController.GetQuestionById)
	beego.InsertFilter("/question/get", beego.BeforeRouter, middleware.AuthMiddleware)
	beego.InsertFilter("/question/get", beego.BeforeRouter, middleware.AdminMiddleware)
	beego.CtrlPost("/question/update", controllers.QuestionController.UpdateQuestion)
	beego.InsertFilter("/question/update", beego.BeforeRouter, middleware.AuthMiddleware)
	beego.InsertFilter("/question/update", beego.BeforeRouter, middleware.AdminMiddleware)
	beego.CtrlPost("/question/list/page", controllers.QuestionController.ListQuestionByPage)
	beego.InsertFilter("/question/list/page", beego.BeforeRouter, middleware.AuthMiddleware)
	beego.InsertFilter("/question/list/page", beego.BeforeRouter, middleware.AdminMiddleware)

	userNs := beego.NewNamespace("/user",
		beego.NSRouter("/register", &controllers.UserController{}, "post:UserRegister"),
		beego.NSRouter("/login", &controllers.UserController{}, "post:UserLogin"),
		beego.NSRouter("/logout", &controllers.UserController{}, "post:UserLogout"),
		beego.NSRouter("/get/login", &controllers.UserController{}, "get:GetLoginUser"),
		beego.NSRouter("/get/vo", &controllers.UserController{}, "get:GetUserVOById"),
		beego.NSRouter("/list/page/vo", &controllers.UserController{}, "get:ListUserVOByPage"),
	)

	questionNs := beego.NewNamespace("/question",
		beego.NSRouter("/get/vo", &controllers.QuestionController{}, "get:GetQuestionVOById"),
		beego.NSRouter("/list/page/vo", &controllers.QuestionController{}, "post:ListQuestionVOByPage"),
	)
	// userNs.Router("/list/page/vo", &controllers.UserController{}, "post:ListUserVOByPage")
	// userNs.Router("/update/my", &controllers.UserController{}, "post:UpdateMyUser").Filter("before", middleware.AuthMiddleware)	//这个filter是作用在userNs上的，不仅仅是当前router

	// questionNs := beego.NewNamespace("/question",
	// 	beego.NSBefore(middleware.AuthMiddleware),  //登录验证中间件
	// 	beego.NSBefore(middleware.AdminMiddleware), //管理员验证中间件
	// 	beego.NSRouter("/add", &controllers.QuestionController{}, "post:AddQuestion"),
	// 	beego.NSRouter("/delete", &controllers.QuestionController{}, "post:DeleteQuestion"),
	// 	beego.NSRouter("/edit", &controllers.QuestionController{}, "post:EditQuestion"),
	// 	beego.NSRouter("/get", &controllers.QuestionController{}, "get:GetQuestionById"),
	// )

	// beego.BeforeRouter：
	// 这个过滤器在路由匹配之后，执行路由处理函数之前执行。
	// 适合用于执行一些前置操作，如身份验证、请求日志记录等。
	// 如果在此过滤器中返回错误，将会终止路由处理函数的执行。

	// beego.BeforeExec：
	// 这个过滤器在路由处理函数执行之后，请求执行之前执行。
	// 适合用于对请求的进一步处理，如请求参数解析、响应头设置等。
	// 如果在此过滤器中返回错误，将会终止请求的执行。

	// 具体区别在于执行的时机，beego.BeforeRouter 在路由匹配后执行，而 beego.BeforeExec 在路由处理函数执行后、请求执行前执行
	beego.AddNamespace(userNs, questionNs)
}
