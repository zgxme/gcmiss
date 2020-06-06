package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["gcmiss/controllers:ArticalController"] = append(beego.GlobalControllerRouter["gcmiss/controllers:ArticalController"],
        beego.ControllerComments{
            Method: "AddArtical",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gcmiss/controllers:ArticalController"] = append(beego.GlobalControllerRouter["gcmiss/controllers:ArticalController"],
        beego.ControllerComments{
            Method: "DeleteArtical",
            Router: `/delete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gcmiss/controllers:ArticalController"] = append(beego.GlobalControllerRouter["gcmiss/controllers:ArticalController"],
        beego.ControllerComments{
            Method: "GetArtical",
            Router: `/get`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gcmiss/controllers:ArticalController"] = append(beego.GlobalControllerRouter["gcmiss/controllers:ArticalController"],
        beego.ControllerComments{
            Method: "GetArticalItem",
            Router: `/item/get`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gcmiss/controllers:ArticalController"] = append(beego.GlobalControllerRouter["gcmiss/controllers:ArticalController"],
        beego.ControllerComments{
            Method: "UpdateArtical",
            Router: `/update`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gcmiss/controllers:CommentController"] = append(beego.GlobalControllerRouter["gcmiss/controllers:CommentController"],
        beego.ControllerComments{
            Method: "AddComment",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gcmiss/controllers:ManagerController"] = append(beego.GlobalControllerRouter["gcmiss/controllers:ManagerController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/delete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gcmiss/controllers:ManagerController"] = append(beego.GlobalControllerRouter["gcmiss/controllers:ManagerController"],
        beego.ControllerComments{
            Method: "DeleteManager",
            Router: `/deleteManager`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gcmiss/controllers:ManagerController"] = append(beego.GlobalControllerRouter["gcmiss/controllers:ManagerController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gcmiss/controllers:PostController"] = append(beego.GlobalControllerRouter["gcmiss/controllers:PostController"],
        beego.ControllerComments{
            Method: "AddPost",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gcmiss/controllers:PostController"] = append(beego.GlobalControllerRouter["gcmiss/controllers:PostController"],
        beego.ControllerComments{
            Method: "DeletePost",
            Router: `/delete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gcmiss/controllers:PostController"] = append(beego.GlobalControllerRouter["gcmiss/controllers:PostController"],
        beego.ControllerComments{
            Method: "GetPost",
            Router: `/get`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gcmiss/controllers:PostController"] = append(beego.GlobalControllerRouter["gcmiss/controllers:PostController"],
        beego.ControllerComments{
            Method: "GetPostItem",
            Router: `/item/get`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gcmiss/controllers:PostController"] = append(beego.GlobalControllerRouter["gcmiss/controllers:PostController"],
        beego.ControllerComments{
            Method: "UpdatePost",
            Router: `/update`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gcmiss/controllers:SessionController"] = append(beego.GlobalControllerRouter["gcmiss/controllers:SessionController"],
        beego.ControllerComments{
            Method: "GetSessionData",
            Router: `/get`,
            AllowHTTPMethods: []string{"get"},
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

    beego.GlobalControllerRouter["gcmiss/controllers:UserController"] = append(beego.GlobalControllerRouter["gcmiss/controllers:UserController"],
        beego.ControllerComments{
            Method: "ActiveUser",
            Router: `/active`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gcmiss/controllers:UserController"] = append(beego.GlobalControllerRouter["gcmiss/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUser",
            Router: `/get`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gcmiss/controllers:UserController"] = append(beego.GlobalControllerRouter["gcmiss/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetProfile",
            Router: `/profile/get`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gcmiss/controllers:UserController"] = append(beego.GlobalControllerRouter["gcmiss/controllers:UserController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gcmiss/controllers:UserController"] = append(beego.GlobalControllerRouter["gcmiss/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateUser",
            Router: `/update`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
