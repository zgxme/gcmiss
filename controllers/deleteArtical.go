/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2020-02-09 12:26:36
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-10 20:53:50
 */
package controllers

import (
	"encoding/json"
	"gcmiss/models"

	"github.com/astaxie/beego"
)

// @Title articalDelete
// @Description user delete	artical
// @Param artical_id	body	string	true	"The id of artical"
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure	2	param error
// @Failure 4009	db error
// @router /delete [post]
func (r *ArticalController) DeleteArtical() {
	articalInfo := make(map[string]interface{})
	defer r.RespData(&r.Resp)
	err := json.Unmarshal(r.Ctx.Input.RequestBody, &articalInfo)
	if err != nil {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
	//TODO check param
	r.Errno = RECODE_OK
	r.Errmsg = RecodeErr(RECODE_OK)
	//获取当前用户
	nickName := r.GetSession("nickname")
	userID := r.GetSession("user_id")
	if userID == nil || nickName == nil {
		r.Errno = AUTH_LOGIN
		r.Errmsg = RecodeErr(AUTH_LOGIN)
		beego.Error(r.errLog(RecodeErr(AUTH_LOGIN)))
		return
	}
	err = models.DeleteArtical(userID.(int64), articalInfo)
	if err != nil {
		r.Errno = DB_ERROR
		r.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
}
