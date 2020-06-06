package controllers

/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2019-12-14 14:24:02
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-23 11:54:26
 */

import (
	"encoding/json"
	"gcmiss/models"
	"strings"
	"unicode/utf8"

	"github.com/astaxie/beego"
)

//if true,check right
func checkSingleParam(param string) bool {
	//check len
	if utf8.RuneCountInString(param) > 20 || utf8.RuneCountInString(param) < 5 {
		return false
	}
	//check space char not exist
	index := strings.Index(param, " ")
	if index != -1 {
		return false
	}
	return true
}

//if true,check right
func checkUserParam(nickName string, passWord string, email string) bool {
	if !checkSingleParam(nickName) || !checkSingleParam(passWord) {
		return false
	}
	if !checkEmailFormat(email) {
		return false
	}

	return true
}

// @Title userRegister
// @Description user regiseter by nickname and password
// @Param	nickname	body	string	true	"The nickname of user register"
// @Param password	body	string	true	"The password of user register"
// @Param email body string true "email of user register"
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure	2	param error
// @Failure 4003	db exist
// @Failure 4009	db error
// @router /register [post]
func (r *UserController) Register() {
	userInfo := make(map[string]interface{})
	r.Errno = RECODE_OK
	// r.Errno = MOCK_TEST
	r.Errmsg = RecodeErr(RECODE_OK)
	defer r.RespData(&r.Resp)
	err := json.Unmarshal(r.Ctx.Input.RequestBody, &userInfo)
	if err != nil {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
	nickName, _ := userInfo["nickname"].(string)
	passWord, _ := userInfo["password"].(string)
	email, _ := userInfo["email"].(string)
	if !checkUserParam(nickName, passWord, email) {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(RecodeErr(PARAM_ERROR)))
		return
	}
	exist, err := models.OneUserExist(nickName)
	if err != nil {
		r.Errno = DB_ERROR
		r.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
	if exist != false {
		r.Errno = DB_EXIST
		r.Errmsg = RecodeErr(DB_EXIST)
		beego.Error(r.errLog(RecodeErr(DB_EXIST)))
		return
	}

	exist, err = models.OneEmailExist(email)
	if err != nil {
		r.Errno = DB_ERROR
		r.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
	if exist != false {
		r.Errno = EMAIL_EXIST
		r.Errmsg = RecodeErr(EMAIL_EXIST)
		beego.Error(r.errLog(RecodeErr(EMAIL_EXIST)))
		return
	}
	//DONE
	//1.put into func of models package
	//2.add create time and update time
	//3.auth check by email (redis?)
	// err = models.RegisterUser(nickName, passWord, email)
	// if err != nil {
	// 	r.Errno = DB_ERROR
	// 	r.Errmsg = RecodeErr(DB_ERROR)
	// 	beego.Error(r.errLog(err.Error()))
	// 	return
	// }
	var auth string
	auth, err = registerSet(nickName, passWord, email)
	if err != nil {
		r.Errno = REDIS_ERROR
		r.Errmsg = RecodeErr(REDIS_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
	err = sendMail(nickName, email, auth)
	if err != nil {
		r.Errno = REDIS_ERROR
		r.Errmsg = RecodeErr(REDIS_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}

}
