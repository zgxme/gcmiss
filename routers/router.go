/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2019-12-13 23:29:09
 * @LastEditors  : Zheng Gaoxiong
 * @LastEditTime : 2020-02-08 02:31:56
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
}
