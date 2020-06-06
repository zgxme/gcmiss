/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2020-04-05 20:05:58
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-23 13:22:01
 */
package controllers

import (
	"gcmiss/models"

	"github.com/astaxie/beego"
)

// @Title profileGet
// @Description get user info
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure 4009	db error
// @router /profile/get [get]
func (r *UserController) GetProfile() {
	defer r.RespData(&r.ProfileResp)
	var userID interface{}
	var err error
	userID, err = r.GetInt64("user_id")
	if err != nil {
		r.ProfileResp.Errno = PARAM_ERROR
		r.ProfileResp.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
	if userID.(int64) <= 0 {
		userID = r.GetSession("user_id")
	}
	if userID == nil {
		r.UserResp.Errno = AUTH_LOGIN
		r.UserResp.Errmsg = RecodeErr(AUTH_LOGIN)
		beego.Error(r.errLog(RecodeErr(AUTH_LOGIN)))
		return
	}
	profileInfo, err := models.GetProfileInfo(userID.(int64))
	if err != nil {
		r.UserResp.Errno = DB_ERROR
		r.UserResp.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
	r.ProfileResp.ProfileInfo = *profileInfo

}
