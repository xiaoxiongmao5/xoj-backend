/*
 * @Author: 小熊 627516430@qq.com
 * @Date: 2023-09-26 10:35:03
 * @LastEditors: 小熊 627516430@qq.com
 * @LastEditTime: 2023-09-27 18:54:53
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
	beego.CtrlPost("/user/register", controllers.UserController.Register)
	beego.CtrlPost("/user/login", controllers.UserController.Login)
	beego.CtrlGet("/user/uinfo", controllers.UserController.GetUserInfoByUserAccount)

	beego.CtrlPost("/question/add", controllers.QuestionController.AddQuestion)
	beego.CtrlPost("/question/delete", controllers.QuestionController.DeleteQuestion)
	beego.CtrlPost("/question/edit", controllers.QuestionController.EditQuestion)
	beego.CtrlGet("/question/get", controllers.QuestionController.GetQuestionById)
	beego.CtrlGet("/question/get/vo", controllers.QuestionController.GetQuestionVOById)
	// 需要管理员
	beego.CtrlPost("/question/update", controllers.QuestionController.UpdateQuestion)

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
