/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2020-05-23 16:36:52
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-23 17:46:40
 */
package controllers

import (
	"gcmiss/models"

	"github.com/astaxie/beego"
)

// @Title articalGet
// @Description user get artical
// @Param post_id int true "The id of post"
// @Param artical_id int true "The id of artical"
// @Param	cursor	int	true	"The cursor of post list"
// @Param limit	int	true	"The limit of post list"
// @Param	desc	int	true	"post list order by desc"
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure	2	param error
// @Failure 4009	db error
// @router /item/get [get]
func (r *ArticalController) GetArticalItem() {
	commentInfo := make(map[string]interface{})
	defer r.RespData(&r.ArticalItemResp)
	var err error
	commentInfo["postId"], err = r.GetInt64("post_id")
	commentInfo["articalId"], err = r.GetInt64("artical_id")
	commentInfo["cursor"], err = r.GetInt64("cursor")
	commentInfo["limit"], err = r.GetInt64("limit")
	commentInfo["desc"], err = r.GetInt("desc")
	if err != nil {
		r.ArticalItemResp.Errno = PARAM_ERROR
		r.ArticalItemResp.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(err.Error()))
	}

	modelErr, post := models.GetArticalItem(commentInfo["articalId"].(int64))
	if modelErr != nil {
		r.ArticalItemResp.Errno = DB_ERROR
		r.ArticalItemResp.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(modelErr.Error()))
		return
	}
	// r.PostResp.HaseMore = true
	// if lenPostList < postInfo["limit"].(int64) {
	// 	r.PostResp.HaseMore = false
	// }
	// Err, commentList := models.GetPostItem(postId,markList)
	//TODO add has more
	commentErr, _, commentList := models.GetCommentList(commentInfo)
	if commentErr != nil {
		r.ArticalItemResp.Errno = DB_ERROR
		r.ArticalItemResp.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(modelErr.Error()))
		return
	}
	imageErr, imageList := models.GetPostImageList(commentInfo)
	if imageErr != nil {
		r.ArticalItemResp.Errno = DB_ERROR
		r.ArticalItemResp.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(modelErr.Error()))
		return
	}
	r.ArticalItemResp.Post = *post
	r.ArticalItemResp.ImageList = *imageList
	r.ArticalItemResp.CommentList = *commentList

	// r.ArticalItemResp.PostList = *postList
}
