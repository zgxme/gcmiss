/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2020-02-09 14:52:23
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-02-09 14:52:47
 */
package controllers

import (
	"gcmiss/models"

	"github.com/astaxie/beego"
)

//GetArticalController get artical controller
type GetArticalController struct {
	BaseController
}

//GetArtical get artical
func (r *GetArticalController) GetArtical() {
	articalInfo := make(map[string]interface{})
	defer r.RespData(&r.ArticalResp)
	var err error
	articalInfo["cursor"], err = r.GetInt64("cursor")
	articalInfo["limit"], err = r.GetInt64("limit")
	articalInfo["desc"], err = r.GetInt("desc")
	if err != nil {
		r.ArticalResp.Errno = PARAM_ERROR
		r.ArticalResp.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(RecodeErr(PARAM_ERROR)))
	}

	modelErr, articalList := models.GetArtical(articalInfo)
	if modelErr != nil {
		r.ArticalResp.Errno = DB_ERROR
		r.ArticalResp.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(modelErr.Error()))
		return
	}
	r.ArticalResp.ArticalList = *articalList
}
