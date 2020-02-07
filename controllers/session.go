/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2019-12-14 21:22:51
 * @LastEditors  : Zheng Gaoxiong
 * @LastEditTime : 2020-02-08 01:27:00
 */
package controllers

//SessionController session controller
type SessionController struct {
	BaseController
}

//GetSessionData get session
func (r *SessionController) GetSessionData() {
	defer r.RespData(&r.Session)
	nickname := r.GetSession("nickname")
	userID := r.GetSession("user_id")
	if nickname != nil {
		r.Session.Nickname = nickname.(string)
		r.Session.UserID = userID.(int64)
	} else {
		r.Errmsg = RecodeErr(AUTH_LOGIN)
		r.Errno = AUTH_LOGIN
	}
}
