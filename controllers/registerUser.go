/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2019-12-14 14:24:02
 * @LastEditors  : Zheng Gaoxiong
 * @LastEditTime : 2020-02-08 00:38:36
 */
package controllers

import (
	"encoding/json"
	"gcmiss/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//RegisterController regiser controller
type RegisterController struct {
	BaseController
}

//Register user register api
func (r *RegisterController) Register() {
	userInfo := make(map[string]interface{})
	defer r.RespData(&r.Resp)
	err := json.Unmarshal(r.Ctx.Input.RequestBody, &userInfo)
	if err != nil {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
	r.Errno = RECODE_OK
	r.Errmsg = RecodeErr(RECODE_OK)
	nickName, _ := userInfo["nickname"].(string)
	passWord, _ := userInfo["password"].(string)
	if nickName == "" || passWord == "" {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(RecodeErr(PARAM_ERROR)))
		return
	}
	exist := models.OneUserExist(nickName)
	o := orm.NewOrm()
	user := models.User{}

	if exist != false {
		r.Errno = DB_EXIST
		r.Errmsg = RecodeErr(DB_EXIST)
		beego.Error(r.errLog(RecodeErr(DB_EXIST)))
		return
	}

	user.Nickname = nickName
	//DONE add salt
	user.Password = GetPassword(passWord)
	profile := models.Profile{}
	user.Profile = &profile
	_, err = o.Insert(&profile)
	if err != nil {
		r.Errno = DB_ERROR
		r.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
	_, err = o.Insert(&user)
	if err != nil {
		r.Errno = DB_ERROR
		r.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
}
