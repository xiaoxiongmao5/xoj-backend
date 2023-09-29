/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-26 10:35:03
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-29 20:13:40
 * @FilePath: /xoj-backend/routers/router.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
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
	"github.com/xiaoxiongmao5/xoj/xoj-backend/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// beego.AutoRouter(&controllers.UserController{})

	beego.CtrlPost("/user/register", controllers.UserController.UserRegister)
	beego.CtrlPost("/user/login", controllers.UserController.UserLogin)
	beego.CtrlPost("/user/logout", controllers.UserController.UserLogout)
	beego.CtrlGet("/user/get/login", controllers.UserController.GetLoginUser)

	// 需要管理员
	beego.CtrlPost("/user/add", controllers.UserController.AddUser)
	beego.CtrlPost("/user/delete", controllers.UserController.DeleteUser)
	beego.CtrlPost("/user/update", controllers.UserController.UpdateUser)
	beego.CtrlGet("/user/get", controllers.UserController.GetUserById)
	beego.CtrlPost("/user/list/page", controllers.UserController.ListUserByPage)
	// 不需要管理员
	beego.CtrlGet("/user/get/vo", controllers.UserController.GetUserVOById)
	beego.CtrlPost("/user/list/page/vo", controllers.UserController.ListUserVOByPage)
	beego.CtrlPost("/user/update/my", controllers.UserController.UpdateMyUser)

	beego.CtrlPost("/question/add", controllers.QuestionController.AddQuestion)
	beego.CtrlPost("/question/delete", controllers.QuestionController.DeleteQuestion)
	beego.CtrlPost("/question/edit", controllers.QuestionController.EditQuestion)
	beego.CtrlGet("/question/get", controllers.QuestionController.GetQuestionById)
	beego.CtrlGet("/question/get/vo", controllers.QuestionController.GetQuestionVOById)
	beego.CtrlPost("/question/list/page/vo", controllers.QuestionController.ListQuestionVOByPage)
	beego.CtrlPost("/question/my/list/page/vo", controllers.QuestionController.ListMyQuestionVOByPage)

	// /question/question_submit/do
	// /question/question_submit/list/page

	// 需要管理员
	beego.CtrlPost("/question/update", controllers.QuestionController.UpdateQuestion)
	beego.CtrlPost("/question/list/page", controllers.QuestionController.ListQuestionByPage)

	// ns := beego.NewNamespace("/v1",
	// 	beego.NSNamespace("/object",
	// 		beego.NSInclude(
	// 			&controllers.ObjectController{},
	// 		),
	// 	),
	// 	beego.NSNamespace("/user",
	// 		beego.NSInclude(
	// 			&controllers.UserController{},
	// 		),
	// 	),
	// )
	// beego.AddNamespace(ns)
}
