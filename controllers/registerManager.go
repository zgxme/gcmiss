/*
 * @Descripttion:	manager register api
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2019-12-15 11:10:28
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-03-21 15:55:18
 */

package controllers

import (
	"encoding/json"
	"gcmiss/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//manager register api
type RegisterManagerController struct {
	SessionController
}

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
func (r *RegisterManagerController) Register() {
	userInfo := make(map[string]interface{})
	defer r.RespData(&r.Resp)
	err := json.Unmarshal(r.Ctx.Input.RequestBody, &userInfo)
	if err != nil {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(r.Errmsg))
		return
	}
	//only manager can register manager
	managerName := userInfo["nickname"].(string)
	nickname := r.GetSession("nickname").(string)
	//add self
	if managerName == nickname {
		r.Errno = DB_ADD_SELF
		r.Errmsg = RecodeErr(DB_ADD_SELF)
		beego.Error(r.errLog(RecodeErr(DB_ADD_SELF)))
		return
	}
	if err != nil {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}

	//获取当前用户id
	userID := models.GetOneuserID(nickname)
	exist := models.OneManagerExist(userID)

	//当前用户不是管理员
	if exist != true || nickname == "" {
		r.Errno = DB_NOT_EXIST
		r.Errmsg = RecodeErr(DB_NOT_EXIST)
		beego.Error(r.errLog(RecodeErr(DB_NOT_EXIST)))
		return
	}

	//添加管理员是否在用户表里
	//不存在用户表中,报错

	userExist := models.OneUserExist(managerName)
	if userExist != true {
		r.Errno = DB_NOT_EXIST
		r.Errmsg = RecodeErr(DB_NOT_EXIST)
		beego.Error(r.errLog(RecodeErr(DB_NOT_EXIST)))
		return
	}
	o := orm.NewOrm()
	userID = models.GetOneuserID(managerName)
	manager := models.Manager{
		UserId: userID,
	}
	_, err = o.Insert(&manager)
	if err != nil {
		r.Errno = DB_ERROR
		r.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
}
