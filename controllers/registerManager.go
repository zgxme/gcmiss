/*
 * @Descripttion:	manager register api
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2019-12-15 11:10:28
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-24 14:32:16
 */

package controllers

import (
	"encoding/json"
	"gcmiss/models"

	"github.com/astaxie/beego"
)

// @Title userRegisterManager
// @Description user regiseter by nickname and password
// @Param	nickname	body	string	true	"The nickname of user register manager"
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure	2	param error
// @Failure 4003	db exist
// @Failure 4004	db not exist
// @Failure 4005	cannot add self
// @Failure 4009	db error
// @router /register [post]
func (r *ManagerController) Register() {
	userInfo := make(map[string]interface{})
	defer r.RespData(&r.Resp)
	err := json.Unmarshal(r.Ctx.Input.RequestBody, &userInfo)
	if err != nil {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
	//only manager can register manager
	managerName := userInfo["nickname"].(string)
	nickname := r.GetSession("nickname").(string)
	//add self
	if !adminAuth(nickname) {
		r.Errno = DB_NOT_EXIST
		r.Errmsg = RecodeErr(DB_NOT_EXIST)
		beego.Error(r.errLog(RecodeErr(DB_NOT_EXIST)))
		return
	}
	if managerName == nickname {
		r.Errno = DB_ADD_SELF
		r.Errmsg = RecodeErr(DB_ADD_SELF)
		beego.Error(r.errLog(RecodeErr(DB_ADD_SELF)))
		return
	}
	if !checkSingleParam(managerName) {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(RecodeErr(PARAM_ERROR)))
		return
	}

	//get now session user id
	userID, err := models.GetOneuserID(nickname)
	exist, err := models.OneManagerExist(userID)
	if err != nil {
		r.Errno = DB_ERROR
		r.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
	//DONE check manager exist
	if exist != true {
		r.Errno = DB_NOT_EXIST
		r.Errmsg = RecodeErr(DB_NOT_EXIST)
		beego.Error(r.errLog(RecodeErr(DB_NOT_EXIST)))
		return
	}

	//now session is not admain
	if exist != true || nickname == "" {
		r.Errno = DB_NOT_EXIST
		r.Errmsg = RecodeErr(DB_NOT_EXIST)
		beego.Error(r.errLog(RecodeErr(DB_NOT_EXIST)))
		return
	}

	//add user must be in user table

	userExist, err := models.OneUserExist(managerName)
	if userExist != true {
		r.Errno = DB_NOT_EXIST
		r.Errmsg = RecodeErr(DB_NOT_EXIST)
		beego.Error(r.errLog(RecodeErr(DB_NOT_EXIST)))
		return
	}
	// o := orm.NewOrm()
	userID, err = models.GetOneuserID(managerName)

	// manager := models.Manager{}
	// _, err = o.Insert(&manager)
	err = models.RegisterManager(userID)
	if err != nil {
		r.Errno = DB_ERROR
		r.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
}
