/*
 * @Descripttion: login api
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2019-12-15 11:22:26
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-24 13:01:19
 */
package controllers

import (
	"encoding/json"
	"fmt"
	"gcmiss/models"
	. "gcmiss/models"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// @Title userLogin
// @Description user login by nickname and passWord
// @Param	nickname	body	string	true	"The nickname of user register"
// @Param passWord	body	string	true	"The passWord of user register"
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
	passWord := userInfo["password"].(string)
	if !checkSingleParam(nickName) || !checkSingleParam(passWord) {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(RecodeErr(PARAM_ERROR)))
		return
	}
	//TODO check user exist
	userExist, err := models.OneUserExist(nickName)
	// fmt.Println(userExist)
	if err != nil {
		r.Errno = DB_ERROR
		r.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
	if userExist != true {
		r.Errno = DB_NOT_EXIST
		r.Errmsg = RecodeErr(DB_NOT_EXIST)
		beego.Error(r.errLog(RecodeErr(DB_NOT_EXIST)))
		return
	}
	o := orm.NewOrm()
	var user User
	// err = o.Raw("SELECT id, nickname,password FROM tb_user WHERE nickname = '?' AND status = 0", nickName).QueryRow(&user)
	var sql string
	//status is not 1
	sql = fmt.Sprintf("SELECT id, nickname, password FROM tb_user WHERE nickname = '%s' AND status = 0", nickName)
	// fmt.Println(sql)
	var terms []orm.Params
	_, err = o.Raw(sql).Values(&terms)
	for _, term := range terms {
		userIDValue := term["id"].(string)
		userID, _ := strconv.ParseInt(userIDValue, 10, 64)
		user = User{Id: userID}
		// fmt.Println("userid", userID)
		err = o.Read(&user)
	}
	// fmt.Println(user)
	if err != nil {
		r.Errno = RECODE_DATAERR
		r.Errmsg = RecodeErr(RECODE_DATAERR)
		beego.Error(r.errLog(RecodeErr(RECODE_DATAERR)))
		return
	}
	// DONE add salt check
	if user.Password != GetPassword(passWord) {
		r.Errno = RECODE_DATAERR
		r.Errmsg = RecodeErr(RECODE_DATAERR)
		beego.Error(r.errLog(RecodeErr(RECODE_DATAERR)))
		return
	}

	//添加Session
	r.SetSession("nickname", userInfo["nickname"])
	r.SetSession("user_id", user.Id)
}
