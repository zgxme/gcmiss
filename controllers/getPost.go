/*
 * @Descripttion: get post list api
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2020-02-08 13:33:10
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-03 10:07:39
 */

package controllers

import (
	"gcmiss/models"

	"github.com/astaxie/beego"
)

// @Title postGet
// @Description user get post
// @Param	cursor	int	true	"The cursor of post list"
// @Param limit	int	true	"The limit of post list"
// @Param	desc	int	true	"post list order by desc"
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure	2	param error
// @Failure 4009	db error
// @router /get [get]
func (r *PostController) GetPost() {
	postInfo := make(map[string]interface{})
	defer r.RespData(&r.PostResp)
	var err error
	postInfo["cursor"], err = r.GetInt64("cursor")
	if err != nil {
		r.PostResp.Errno = PARAM_ERROR
		r.PostResp.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(err.Error()))
	}
	postInfo["limit"], err = r.GetInt64("limit")
	if err != nil {
		r.PostResp.Errno = PARAM_ERROR
		r.PostResp.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(err.Error()))
	}
	postInfo["desc"], err = r.GetInt("desc")
	if err != nil {
		r.PostResp.Errno = PARAM_ERROR
		r.PostResp.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(err.Error()))
	}
	postInfo["tag"], err = r.GetInt("tag")
	if err != nil {
		r.PostResp.Errno = PARAM_ERROR
		r.PostResp.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(err.Error()))
	}

	modelErr, lenPostList, postList := models.GetPost(postInfo)
	if modelErr != nil {
		r.PostResp.Errno = DB_ERROR
		r.PostResp.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(modelErr.Error()))
		return
	}
	r.PostResp.HaseMore = true
	if lenPostList < postInfo["limit"].(int64) {
		r.PostResp.HaseMore = false
	}
	r.PostResp.PostList = *postList
}
