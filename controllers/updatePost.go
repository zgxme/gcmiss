/*
 * @Descripttion: user update post api
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2020-02-08 01:53:16
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-10 21:16:05
 */
package controllers

import (
	"encoding/json"
	"gcmiss/models"

	"github.com/astaxie/beego"
)

//update post api
type UpdatePostController struct {
	BaseController
}

// @Title postUpdate
// @Description user update post info
// @Param post_id	formData	string	false	"The id of post"
// @Param	title	formData	string	true	"The	title of post"
// @Param content formData	string	true	"The content of post"
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure	2	param error
// @Failure 4002	db data error
// @Failure 4009	db error
// @Failure 4012	get file error
// @Failure	4013	file format error
// @router /update [post]
func (r *PostController) UpdatePost() {
	postInfo := make(map[string]interface{})
	defer r.RespData(&r.Resp)
	err := json.Unmarshal(r.Ctx.Input.RequestBody, &postInfo)
	if err != nil {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
	if !checkContent(postInfo["content"].(string)) {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
	if !checkTitle(postInfo["title"].(string)) {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
	postID := int64(postInfo["post_id"].(float64))
	if postID <= 0 {
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
	err = models.UpdatePost(userID.(int64), postInfo)
	if err != nil {
		r.Errno = DB_ERROR
		r.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
}
