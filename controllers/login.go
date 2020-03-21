/*
 * @Descripttion: login api
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2019-12-15 11:22:26
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-03-21 14:31:42
 */
package controllers

import (
	"encoding/json"
	"gcmiss/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// @Title userLogin
// @Description user login by nickname and password
// @Param	nickname	body	string	true	"The nickname of user register"
// @Param password	body	string	true	"The password of user register"
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure	2	param error
// @Failure 4002	db data error
// @router /login [post]
func (r *SessionController) Login() {
	defer r.RespData(&r.Resp)
	userInfo := make(map[string]interface{})
	err := json.Unmarshal(r.Ctx.Input.RequestBody, &userInfo)
	if err != nil {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
	nickName := userInfo["nickname"].(string)
	password := userInfo["password"].(string)
	if nickName == "" || password == "" {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(RecodeErr(PARAM_ERROR)))
		return
	}
	o := orm.NewOrm()
	user := models.User{Nickname: nickName}
	qs := o.QueryTable("tb_user")
	err = qs.Filter("nickname", nickName).One(&user)
	if err != nil {
		r.Errno = RECODE_DATAERR
		r.Errmsg = RecodeErr(RECODE_DATAERR)
		beego.Error(r.errLog(RecodeErr(RECODE_DATAERR)))
		return
	}
	//DONE add salt check
	if user.Password != GetPassword(password) {
		r.Errno = RECODE_DATAERR
		r.Errmsg = RecodeErr(RECODE_DATAERR)
		beego.Error(r.errLog(RecodeErr(RECODE_DATAERR)))
		return
	}
	//添加Session
	r.SetSession("nickname", userInfo["nickname"])
	r.SetSession("user_id", user.Id)
}
