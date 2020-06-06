/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2020-03-28 16:27:00
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-24 10:25:57
 */

package controllers

import (
	"gcmiss/models"

	"github.com/astaxie/beego"
)

// @Title userGet
// @Description get user info
// @Param user_id	int	false	"The id of user"
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure 4009	db error
// @router /get [get]
func (r *UserController) GetUser() {
	defer r.RespData(&r.UserResp)
	var userID interface{}
	var err error
	userID, err = r.GetInt64("user_id")
	if err != nil {
		r.UserResp.Errno = PARAM_ERROR
		r.UserResp.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
	if userID.(int64) <= 0 {
		userID = r.GetSession("user_id")
	}
	// fmt.Println("userID is", userID)
	if userID == nil {
		r.UserResp.Errno = AUTH_LOGIN
		r.UserResp.Errmsg = RecodeErr(AUTH_LOGIN)
		beego.Error(r.errLog(RecodeErr(AUTH_LOGIN)))
		return
	}
	userInfo, err := models.GetUserInfo(userID.(int64))
	if err != nil {
		r.UserResp.Errno = DB_ERROR
		r.UserResp.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
	r.UserResp.UserInfo = *userInfo

}
