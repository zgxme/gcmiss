/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2019-12-14 15:15:04
 * @LastEditors  : Zheng Gaoxiong
 * @LastEditTime : 2020-02-08 18:52:29
 */
package controllers

import (
	"fmt"
	"gcmiss/models"

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
	Resp
	Nickname string `json:"nickname"`
	UserID   int64  `json:"user_id"`
}

//PostResp response
type PostResp struct {
	Resp
	PostList []models.PostItem `json:"post_list"`
}

//DataBase request data
type DataBase struct {
	Resp
	Session
	PostResp
}

//BaseController base controller
type BaseController struct {
	beego.Controller
	DataBase
}

//Prepare 1.overridePrepare method, add requestId 2.init Resp
func (c *BaseController) Prepare() {
	RequestID := UniqueId()
	c.Ctx.Input.SetData("requestId", RequestID)
	c.baseLog()
	c.RequestID = RequestID
	c.Session.RequestID = RequestID
	c.PostResp.RequestID = RequestID
	c.Errno = RECODE_OK
	c.Session.Errno = RECODE_OK
	c.PostResp.Errno = RECODE_OK
	c.Errmsg = recodeText[RECODE_OK]
	c.Session.Errmsg = recodeText[RECODE_OK]
	c.PostResp.Errmsg = recodeText[RECODE_OK]
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
