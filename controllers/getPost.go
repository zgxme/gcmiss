/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2020-02-08 13:33:10
 * @LastEditors  : Zheng Gaoxiong
 * @LastEditTime : 2020-02-08 20:02:41
 */

package controllers

import (
	"gcmiss/models"

	"github.com/astaxie/beego"
)

//GetPostController get post controller
type GetPostController struct {
	BaseController
}

//GetPost get post
func (r *GetPostController) GetPost() {
	postInfo := make(map[string]interface{})
	defer r.RespData(&r.PostResp)
	var err error
	postInfo["cursor"], err = r.GetInt64("cursor")
	postInfo["limit"], err = r.GetInt64("limit")
	postInfo["desc"], err = r.GetInt("desc")
	if err != nil {
		r.PostResp.Errno = PARAM_ERROR
		r.PostResp.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(RecodeErr(PARAM_ERROR)))
	}

	modelErr, postList := models.GetPost(postInfo)
	if modelErr != nil {
		r.PostResp.Errno = DB_ERROR
		r.PostResp.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(modelErr.Error()))
		return
	}
	r.PostResp.PostList = *postList
}
