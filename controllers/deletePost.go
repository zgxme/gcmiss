/*
 * @Descripttion:delete post api
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2020-02-03 23:15:01
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-10 21:13:16
 */
package controllers

import (
	"encoding/json"
	"gcmiss/models"

	"github.com/astaxie/beego"
)

// @Title postDelete
// @Description user delete post
// @Param post_id	body	string	true	"The id of post"
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure	2	param error
// @Failure 4009	db error
// @router /delete [post]
func (r *PostController) DeletePost() {
	postInfo := make(map[string]interface{})
	defer r.RespData(&r.Resp)
	err := json.Unmarshal(r.Ctx.Input.RequestBody, &postInfo)
	if err != nil {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(err.Error()))
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
	err = models.DeletePost(userID.(int64), postInfo)
	if err != nil {
		r.Errno = DB_ERROR
		r.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
}
