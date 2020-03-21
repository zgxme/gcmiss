package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["gcmiss/controllers:AddPostController"] = append(beego.GlobalControllerRouter["gcmiss/controllers:AddPostController"],
        beego.ControllerComments{
            Method: "AddPost",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gcmiss/controllers:DeletePostController"] = append(beego.GlobalControllerRouter["gcmiss/controllers:DeletePostController"],
        beego.ControllerComments{
            Method: "DeletePost",
            Router: `/delete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gcmiss/controllers:RegisterController"] = append(beego.GlobalControllerRouter["gcmiss/controllers:RegisterController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gcmiss/controllers:RegisterManagerController"] = append(beego.GlobalControllerRouter["gcmiss/controllers:RegisterManagerController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gcmiss/controllers:SessionController"] = append(beego.GlobalControllerRouter["gcmiss/controllers:SessionController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gcmiss/controllers:SessionController"] = append(beego.GlobalControllerRouter["gcmiss/controllers:SessionController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gcmiss/controllers:UpdateUserController"] = append(beego.GlobalControllerRouter["gcmiss/controllers:UpdateUserController"],
        beego.ControllerComments{
            Method: "UpdateUser",
            Router: `/update`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
