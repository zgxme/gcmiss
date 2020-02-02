package main

import (
	logs "gcmiss/logs"
	_ "gcmiss/models"
	modles "gcmiss/models"
	_ "gcmiss/routers"
	session "gcmiss/session"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/astaxie/beego/session/redis"
)

func init() {
	modles.Init()
	logs.Init()
	session.Init()

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"http://127.0.0.1:8088", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
}

func main() {
	beego.SetStaticPath("css", "static/css")
	beego.SetStaticPath("js", "static/js")
	beego.SetStaticPath("fonts", "static/fonts")
	beego.SetStaticPath("img", "static/img")
	beego.Run()
}
