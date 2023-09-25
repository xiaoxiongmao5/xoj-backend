package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/xiaoxiongmao5/xoj/xoj-backend/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/xiaoxiongmao5/xoj/xoj-backend/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiaoxiongmao5/xoj/xoj-backend/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/xiaoxiongmao5/xoj/xoj-backend/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiaoxiongmao5/xoj/xoj-backend/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/xiaoxiongmao5/xoj/xoj-backend/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiaoxiongmao5/xoj/xoj-backend/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/xiaoxiongmao5/xoj/xoj-backend/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiaoxiongmao5/xoj/xoj-backend/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/xiaoxiongmao5/xoj/xoj-backend/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiaoxiongmao5/xoj/xoj-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xiaoxiongmao5/xoj/xoj-backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiaoxiongmao5/xoj/xoj-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xiaoxiongmao5/xoj/xoj-backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiaoxiongmao5/xoj/xoj-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xiaoxiongmao5/xoj/xoj-backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiaoxiongmao5/xoj/xoj-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xiaoxiongmao5/xoj/xoj-backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiaoxiongmao5/xoj/xoj-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xiaoxiongmao5/xoj/xoj-backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiaoxiongmao5/xoj/xoj-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xiaoxiongmao5/xoj/xoj-backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiaoxiongmao5/xoj/xoj-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xiaoxiongmao5/xoj/xoj-backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/user/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiaoxiongmao5/xoj/xoj-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xiaoxiongmao5/xoj/xoj-backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/user/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiaoxiongmao5/xoj/xoj-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xiaoxiongmao5/xoj/xoj-backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUserInfoByUserAccount",
            Router: `/user/uinfo`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
