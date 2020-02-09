/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2019-12-13 23:29:09
 * @LastEditors  : Zheng Gaoxiong
 * @LastEditTime : 2020-02-09 13:19:12
 */
package routers

import (
	"gcmiss/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//api
	beego.Router("*", &controllers.MainController{})
	beego.Router("/api/v1/user/register", &controllers.RegisterController{}, "post:Register")
	beego.Router("/api/v1/login", &controllers.SessionController{}, "post:Login")
	beego.Router("/api/v1/logout", &controllers.SessionController{}, "post:Logout")
	beego.Router("/api/v1/session/get", &controllers.SessionController{}, "get:GetSessionData")
	beego.Router("/api/v1/manager/register", &controllers.RegisterManagerController{}, "post:Register")
	beego.Router("/api/v1/user/update", &controllers.UpdateUserController{}, "post:UpdateUser")
	beego.Router("/api/v1/post/add", &controllers.AddPostController{}, "post:AddPost")
	beego.Router("/api/v1/post/delete", &controllers.DeletePostController{}, "post:DeletePost")
	beego.Router("/api/v1/post/update", &controllers.UpdatePostController{}, "post:UpdatePost")
	beego.Router("/api/v1/post/getall", &controllers.GetPostController{}, "get:GetPost")
	beego.Router("/api/v1/artical/add", &controllers.AddArticalController{}, "post:AddArtical")
	beego.Router("/api/v1/artical/delete", &controllers.DeleteArticalController{}, "post:DeleteArtical")
	beego.Router("/api/v1/artical/update", &controllers.UpdateArticalController{}, "post:UpdateArtical")
	beego.Router("/api/v1/artical/getall", &controllers.GetArticalController{}, "get:GetArtical")
}
