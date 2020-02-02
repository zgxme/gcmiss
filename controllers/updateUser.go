/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2019-12-15 21:42:55
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-02-03 00:53:04
 */
/**
 * @Author: zhenggaoxiong
 * @Description: update user info
 * @File:  updateUser
 * @Date: 2019/12/15 21:42
 */
package controllers

import (
	"encoding/json"
	"gcmiss/models"

	"github.com/astaxie/beego"
)

type UpdateUserController struct {
	SessionController
	models.User
}

//TODO user nickname update

func (r *UpdateUserController) UpdateUser() {
	userInfo := make(map[string]interface{})
	defer r.RespData(&r.Resp)
	r.Errno = RECODE_OK
	r.Errmsg = RecodeErr(RECODE_OK)
	err := json.Unmarshal(r.Ctx.Input.RequestBody, &userInfo)
	if err != nil {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
	//获取当前用户
	nickName := r.GetSession("nickname")
	userID := r.GetSession("user_id")
	if userID == nil || nickName == nil {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(RecodeErr(PARAM_ERROR)))
		return
	}
	//newNickname := userInfo["new_nickname"].(string)
	//exist := models.OneUserExist(newNickname)
	//if exist {
	//	r.Errno = DB_EXIST
	//	r.Errmsg = RecodeErr(DB_EXIST)
	//	beego.Error(u.errLog(RecodeErr(DB_EXIST)))
	//	return
	//}

	//if newNickname == ""{
	//	r.Errno = PARAM_ERROR
	//	r.Errmsg = RecodeErr(PARAM_ERROR)
	//	beego.Error(u.errLog(RecodeErr(PARAM_ERROR)))
	//	return
	//}
	err = models.UpdateProfile(userID, userInfo)
	if err != nil {
		r.Errno = DB_ERROR
		r.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
}
