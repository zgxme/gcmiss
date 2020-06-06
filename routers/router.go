/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2019-12-13 23:29:09
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-01 11:55:32
 */
// @APIVersion 1.0.0
// @Title gcmiss API
// @Description The API of gcmiss
// @Contact purifiedzheng@gmail.com
package routers

import (
	"gcmiss/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//api
	ns := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
				&controllers.SessionController{},
			),
		),
		beego.NSNamespace("/manager",
			beego.NSInclude(
				&controllers.ManagerController{},
			),
		),
		beego.NSNamespace("/post",
			beego.NSInclude(
				&controllers.PostController{},
			),
		),
		beego.NSNamespace("/session",
			beego.NSInclude(
				&controllers.SessionController{},
			),
		),
		beego.NSNamespace("/artical",
			beego.NSInclude(
				&controllers.ArticalController{},
			),
		),
		beego.NSNamespace("comment",
			beego.NSInclude(
				&controllers.CommentController{},
			),
		),
		// beego.NSNamespace("/manager",
		// 	beego.NSInclude(
		// 		&controllers.RegisterManagerController{},
		// 	),
		// ),
		// beego.NSNamespace("/session",
		// 	beego.NSInclude(
		// 		&controllers.SessionController{},
		// 	),
		// ),
		// beego.NSNamespace("post",
		// 	beego.NSInclude(
		// 		&controllers.AddPostController{},
		// 		&controllers.DeletePostController{},
		// 		&controllers.UpdatePostController{},
		// 		&controllers.GetPostController{},
		// 	),
		// ),
		// beego.NSNamespace("artical",
		// 	beego.NSInclude(
		// 		&controllers.AddArticalController{},
		// 		&controllers.DeleteArticalController{},
		// 		&controllers.UpdateArticalController{},
		// 		&controllers.GetArticalController{},
		// 	),
		// ),
	)
	beego.Router("*", &controllers.MainController{})
	beego.AddNamespace(ns)
	beego.SetStaticPath("/swagger", "swagger")
}
