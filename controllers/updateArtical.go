/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2020-02-09 13:04:03
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-10 10:03:03
 */
package controllers

import (
	"encoding/json"
	"gcmiss/models"

	"github.com/astaxie/beego"
)

// @Title	articalUpdate
// @Description user update artical info
// @Param artical_id	formData	string	false	"The id of artical"
// @Param	artical_name	formData	string	true	"The	title of artical"
// @Param artical_desc formData	string	true	"The content of artical"
// @Param artical_price formData	string	true	"The content of artical"
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure	2	param error
// @Failure 4002	db data error
// @Failure 4009	db error
// @Failure 4012	get file error
// @Failure	4013	file format error
// @router /update [post]
func (r *ArticalController) UpdateArtical() {
	articalInfo := make(map[string]interface{})
	defer r.RespData(&r.Resp)
	err := json.Unmarshal(r.Ctx.Input.RequestBody, &articalInfo)
	if err != nil {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
	if int64(articalInfo["artical_id"].(float64)) <= 0 {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.Errmsg)
		return
	}
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
	err = models.UpdateArtical(userID.(int64), articalInfo)
	if err != nil {
		r.Errno = DB_ERROR
		r.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
}
