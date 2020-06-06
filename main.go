/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2019-12-13 23:29:09
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-04 14:00:09
 */
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
		AllowOrigins:     []string{"http://127.0.0.1:8088", "http://192.168.1.9:8080", "http://192.168.1.9:8088", "http://192.168.1.9:8081"},
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
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
