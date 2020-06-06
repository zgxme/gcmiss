/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2020-04-04 13:40:38
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-23 18:28:55
 */
package controllers

import (
	"encoding/json"
	"gcmiss/models"
	"strings"
	"unicode/utf8"

	"github.com/astaxie/beego"
)

//if true,check right
func checkSingleComment(param string) bool {
	//check len
	if utf8.RuneCountInString(param) > 50 || utf8.RuneCountInString(param) < 5 {
		return false
	}
	//check space char not exist
	index := strings.Index(param, " ")
	if index != -1 {
		return false
	}
	return true
}

// @Title commentAdd
// @Description user add comment
// @Param user_to	body	string	true	"user_id of comment to"
// @Param comment	body	string	true	"the content of comment"
// @Param post_id	body	string	true	"post_id of comment"
// @Param artical_id	body	string	true	"artical_id of comment"
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure	2	param error
// @Failure 4009	db error
// @router /add [post]
func (r *CommentController) AddComment() {
	defer r.RespData(&r.Resp)
	commentInfo := make(map[string]interface{})
	defer r.RespData(&r.Resp)
	r.Errno = RECODE_OK
	r.Errmsg = RecodeErr(RECODE_OK)
	err := json.Unmarshal(r.Ctx.Input.RequestBody, &commentInfo)
	if err != nil {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
	nickName := r.GetSession("nickname")
	userID := r.GetSession("user_id")
	if userID == nil || nickName == nil {
		r.Errno = AUTH_LOGIN
		r.Errmsg = RecodeErr(AUTH_LOGIN)
		beego.Error(r.errLog(RecodeErr(AUTH_LOGIN)))
		return
	}

	Comment := commentInfo["comment"].(string)
	if !checkSingleComment(Comment) {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(RecodeErr(PARAM_ERROR)))
		return
	}
	UserTo := int64(commentInfo["user_to"].(float64))
	PostID := int64(commentInfo["post_id"].(float64))
	err, postItem := models.GetPostItem(PostID)
	if err != nil {
		r.Errno = DB_ERROR
		r.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
	//TODO check articalID
	ArticalID := int64(commentInfo["artical_id"].(float64))
	err, articalItem := models.GetArticalItem(ArticalID)
	if err != nil {
		r.Errno = DB_ERROR
		r.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
	if articalItem.ArticalID == 0 && postItem.PosterId == 0 {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(r.Errmsg))
		return
	}
	if err := models.AddComment(userID.(int64), UserTo, Comment, PostID, ArticalID); err != nil {
		r.Errno = DB_ERROR
		r.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}

}
