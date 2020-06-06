/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2020-05-05 14:58:47
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-23 22:31:06
 */
package controllers

import (
	"encoding/json"
	"gcmiss/models"

	"github.com/astaxie/beego"
)

// @Title userActive
// @Description active user info
// @Param auth string	true	"The auth of user"
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure 4009	db error
// @router /active [post]
func (r *UserController) ActiveUser() {
	authInfo := make(map[string]interface{})
	defer r.RespData(&r.Resp)
	err := json.Unmarshal(r.Ctx.Input.RequestBody, &authInfo)
	if err != nil {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(r.Errmsg))
		return
	}
	auth := authInfo["auth"].(string)
	infoList, err := checkRegister(auth)
	// fmt.Println(infoList)
	if err != nil {
		r.Errno = REDIS_ERROR
		r.Errmsg = RecodeErr(REDIS_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
	var UserID int64
	isAdmin := adminAuth(infoList[0])
	UserID, err = models.RegisterUser(infoList[0], infoList[1], infoList[2], isAdmin)
	if err != nil {
		r.Errno = DB_ERROR
		r.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
	r.SetSession("nickname", infoList[0])
	r.SetSession("user_id", UserID)
}
