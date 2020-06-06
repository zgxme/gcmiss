/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2020-05-03 08:23:48
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-24 11:19:21
 */
package controllers

import (
	"encoding/json"
	"gcmiss/models"

	"github.com/astaxie/beego"
)

// @Title ManagerDeleteManager
// @Description manager delete manager
// @Param	nickname	body	string	true	"The nickname of manager name"
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure	2	param error
// @Failure 4003	db exist
// @Failure 4004	db not exist
// @Failure 4006	cannot delete self
// @Failure 4009	db error
// @router /deleteManager [post]
func (r *ManagerController) DeleteManager() {
	userInfo := make(map[string]interface{})
	defer r.RespData(&r.Resp)
	err := json.Unmarshal(r.Ctx.Input.RequestBody, &userInfo)
	if err != nil {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(r.Errmsg))
		return
	}
	//only manager can delete manager
	managerName := userInfo["nickname"].(string)
	nickname := r.GetSession("nickname").(string)
	//delete self
	if managerName == nickname {
		r.Errno = DB_DELETE_SELF
		r.Errmsg = RecodeErr(DB_DELETE_SELF)
		beego.Error(r.errLog(RecodeErr(DB_DELETE_SELF)))
		return
	}
	if !adminAuth(nickname) {
		r.Errno = DB_NOT_EXIST
		r.Errmsg = RecodeErr(DB_NOT_EXIST)
		beego.Error(r.errLog(RecodeErr(DB_NOT_EXIST)))
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

	//now session is not admain
	if exist != true || nickname == "" {
		r.Errno = DB_NOT_EXIST
		r.Errmsg = RecodeErr(DB_NOT_EXIST)
		beego.Error(r.errLog(RecodeErr(DB_NOT_EXIST)))
		return
	}

	//add user must be in user table
	//TODO fix error
	userExist, err := models.OneUserExist(managerName)
	if userExist != true {
		r.Errno = DB_NOT_EXIST
		r.Errmsg = RecodeErr(DB_NOT_EXIST)
		beego.Error(r.errLog(RecodeErr(DB_NOT_EXIST)))
		return
	}
	userID, err = models.GetOneuserID(managerName)
	err = models.DeleteManager(userID)
	if err != nil {
		r.Errno = DB_ERROR
		r.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
}
