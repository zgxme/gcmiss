/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2020-02-09 14:52:23
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-23 15:51:55
 */
package controllers

import (
	"gcmiss/models"

	"github.com/astaxie/beego"
)

// @Title articalGet
// @Description user get artical
// @Param	cursor	int	true	"The cursor of artical list"
// @Param limit	int	true	"The limit of artical list"
// @Param	desc	int	true	"artical list order by desc"
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure	2	param error
// @Failure 4009	db error
// @router /get [get]
func (r *ArticalController) GetArtical() {
	//TODO check
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

	modelErr, lenArticalList, articalList := models.GetArtical(articalInfo)
	if modelErr != nil {
		r.ArticalResp.Errno = DB_ERROR
		r.ArticalResp.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(modelErr.Error()))
		return
	}
	r.ArticalResp.HaseMore = true
	if lenArticalList < articalInfo["limit"].(int64) {
		r.ArticalResp.HaseMore = false
	}
	r.ArticalResp.ArticalList = *articalList
}
