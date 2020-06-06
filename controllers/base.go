/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2019-12-14 15:15:04
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-23 17:04:16
 */
package controllers

import (
	"fmt"
	. "gcmiss/models"

	"github.com/astaxie/beego"
)

//Resp respose base
type Resp struct {
	Errmsg    string `json:"errmsg"`
	Errno     int    `json:"errno"`
	RequestID string `json:"request_id"`
}

//Session response
type Session struct {
	Nickname string `json:"nickname"`
	UserID   int64  `json:"user_id"`
	Resp
}

//PostResp response
type PostResp struct {
	PostList []PostItem `json:"post_list"`
	HaseMore bool       `json:"has_more"`
	Resp
}

//PostItem response
type PostItemResp struct {
	Post        PostItem      `json:"post"`
	CommentList []CommentItem `json:"comment_list"`
	ImageList   []ImageItem   `json:"image_list"`
	Resp
}

//ArticalResp response
type ArticalResp struct {
	ArticalList []ArticalItem `json:"post_list"`
	HaseMore    bool          `json:"has_more"`
	Resp
}

//ArticalItem response
type ArticalItemResp struct {
	Post        ArticalItem   `json:"post"`
	CommentList []CommentItem `json:"comment_list"`
	ImageList   []ImageItem   `json:"image_list"`
	Resp
}

//UserResp response
type UserResp struct {
	UserInfo UserItem `json:"user_info"`
	Resp
}

//ProfileResp response
type ProfileResp struct {
	ProfileInfo ProfileItemResp `json:"profile_info"`
	Resp
}

//DataBase request data
type DataBase struct {
	Resp
	Session
	PostItemResp
	ArticalItemResp
	PostResp
	ArticalResp
	UserResp
	ProfileResp
}

//BaseController base controller
type BaseController struct {
	beego.Controller
	DataBase
	endpoint string
	buckname string
}

//user api
type UserController struct {
	BaseController
}

//manager api
type ManagerController struct {
	SessionController
}

//post api
type PostController struct {
	BaseController
}

//artical api
type ArticalController struct {
	BaseController
}

//session api
type SessionController struct {
	BaseController
}

//comment api
type CommentController struct {
	BaseController
}

//Prepare 1.overridePrepare method, add requestId 2.init Resp
func (c *BaseController) Prepare() {
	c.endpoint = beego.AppConfig.String("endpoint")
	c.buckname = beego.AppConfig.String("buckname")
	RequestID := UniqueId()
	c.Ctx.Input.SetData("requestId", RequestID)
	c.baseLog()
	c.RequestID = RequestID
	c.Errno = RECODE_OK
	c.Errmsg = recodeText[RECODE_OK]

	c.Session.RequestID = RequestID
	c.Session.Errno = RECODE_OK
	c.Session.Errmsg = recodeText[RECODE_OK]

	c.PostResp.RequestID = RequestID
	c.PostResp.Errno = RECODE_OK
	c.PostResp.Errmsg = recodeText[RECODE_OK]

	c.ArticalResp.RequestID = RequestID
	c.ArticalResp.Errno = RECODE_OK
	c.ArticalResp.Errmsg = recodeText[RECODE_OK]

	c.UserResp.RequestID = RequestID
	c.UserResp.Errno = RECODE_OK
	c.UserResp.Errmsg = recodeText[RECODE_OK]

	c.PostItemResp.RequestID = RequestID
	c.PostItemResp.Errno = RECODE_OK
	c.PostItemResp.Errmsg = recodeText[RECODE_OK]

	c.ProfileResp.RequestID = RequestID
	c.ProfileResp.Errno = RECODE_OK
	c.ProfileResp.Errmsg = recodeText[RECODE_OK]
}

//GetreqID get request id
func (c *BaseController) GetreqID() string {
	return c.Ctx.Input.GetData("requestId").(string)
}

//baseLog access Log
func (c *BaseController) baseLog() {
	logID := c.GetreqID()
	traceID := c.GetreqID()
	protocal := c.Ctx.Input.Protocol()
	uri := c.Ctx.Input.URI()
	method := c.Ctx.Input.Method()
	ip := c.Ctx.Input.IP()
	userAgent := c.Ctx.Input.UserAgent()
	logStr := fmt.Sprintf("logID:[%s] traceID:[%s] ip:[%s] protocal:[%s] method:[%s] uri:[%s] userAgent:[%s]", logID, traceID, ip, protocal, method, uri, userAgent)
	beego.Info(logStr)
}

//errLog err log
func (c *BaseController) errLog(err string) string {
	logID := c.GetreqID()
	traceID := c.GetreqID()
	errLog := fmt.Sprintf("logID:[%s] traceID:[%s] error:[%s]", logID, traceID, err)
	return errLog
}

//RespData translate *RespData to json
func (c *BaseController) RespData(resp interface{}) {
	c.Data["json"] = &resp
	c.ServeJSON()
}

//RespDataV2 translate
//TODO default value
func (c *BaseController) RespDataV2(resp map[string]interface{}) {
	c.Data["json"] = resp
	c.ServeJSON()
}
