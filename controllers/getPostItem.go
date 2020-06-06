/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2020-04-04 15:31:22
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-23 17:12:36
 */
package controllers

import (
	"gcmiss/models"

	"github.com/astaxie/beego"
)

// @Title postGet
// @Description user get post
// @Param post_id int true "The id of post"
// @Param artical_id int true "The id of artical"
// @Param	cursor	int	true	"The cursor of post list"
// @Param limit	int	true	"The limit of post list"
// @Param	desc	int	true	"post list order by desc"
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure	2	param error
// @Failure 4009	db error
// @router /item/get [get]
func (r *PostController) GetPostItem() {
	commentInfo := make(map[string]interface{})
	defer r.RespData(&r.PostItemResp)
	var err error
	commentInfo["postId"], err = r.GetInt64("post_id")
	commentInfo["articalId"], err = r.GetInt64("artical_id")
	commentInfo["cursor"], err = r.GetInt64("cursor")
	commentInfo["limit"], err = r.GetInt64("limit")
	commentInfo["desc"], err = r.GetInt("desc")
	if err != nil {
		r.PostItemResp.Errno = PARAM_ERROR
		r.PostItemResp.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(err.Error()))
	}

	modelErr, post := models.GetPostItem(commentInfo["postId"].(int64))
	if modelErr != nil {
		r.PostItemResp.Errno = DB_ERROR
		r.PostItemResp.Errmsg = RecodeErr(DB_ERROR)
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
		r.PostItemResp.Errno = DB_ERROR
		r.PostItemResp.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(modelErr.Error()))
		return
	}
	imageErr, imageList := models.GetPostImageList(commentInfo)
	if imageErr != nil {
		r.PostItemResp.Errno = DB_ERROR
		r.PostItemResp.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(modelErr.Error()))
		return
	}
	r.PostItemResp.Post = *post
	r.PostItemResp.ImageList = *imageList
	r.PostItemResp.CommentList = *commentList

	// r.PostItemResp.PostList = *postList
}
