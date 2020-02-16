/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2020-02-09 13:04:03
 * @LastEditors  : Zheng Gaoxiong
 * @LastEditTime : 2020-02-16 11:11:00
 */
package controllers

import (
	"encoding/json"
	"gcmiss/models"

	"github.com/astaxie/beego"
)

//update artical controller
type UpdateArticalController struct {
	BaseController
}

//UpdateArtical update artical
func (r *UpdateArticalController) UpdateArtical() {
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
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(RecodeErr(PARAM_ERROR)))
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
