/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2019-12-13 23:29:09
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-02-16 16:17:33
 */
package controllers

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}
