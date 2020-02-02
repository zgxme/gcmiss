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
	beego.Router("/api/v1/user/update", &controllers.UpdateUserController{}, "post:Update")
}
