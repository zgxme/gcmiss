/*
 * @Descripttion: session api
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2019-12-14 21:22:51
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-03-21 14:35:33
 */
package controllers

//session api
type SessionController struct {
	BaseController
}

// @Title getSession
// @Description user get session
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure 5004 auth error
// @router /get [get]

func (r *SessionController) GetSessionData() {
	defer r.RespData(&r.Session)
	nickname := r.GetSession("nickname")
	userID := r.GetSession("user_id")
	if nickname != nil {
		r.Session.Nickname = nickname.(string)
		r.Session.UserID = userID.(int64)
	} else {
		r.Session.Errmsg = RecodeErr(AUTH_LOGIN)
		r.Session.Errno = AUTH_LOGIN
	}
}
