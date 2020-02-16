/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2020-02-09 11:14:43
 * @LastEditors  : Zheng Gaoxiong
 * @LastEditTime : 2020-02-16 11:56:13
 */
package controllers

import (
	"encoding/json"
	"gcmiss/models"

	"github.com/astaxie/beego"
)

//AddArticalController add artical controller
type AddArticalController struct {
	BaseController
}

//AddArtical add artical
func (r *AddArticalController) AddArtical() {
	articalInfo := make(map[string]interface{})
	defer r.RespData(&r.Resp)
	err := json.Unmarshal(r.Ctx.Input.RequestBody, &articalInfo)
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
	err = models.AddArtical(userID.(int64), articalInfo)
	if err != nil {
		r.Errno = DB_ERROR
		r.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
}
