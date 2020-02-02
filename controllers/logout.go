/**
 * @Author: zhenggaoxiong
 * @Description: logout api
 * @File:  logout
 * @Date: 2019/12/15 15:14
 */

package controllers

//Logout logout api
func (r *SessionController) Logout() {
	// resp := make(map[string]interface{})
	// reqID := r.GetreqID()
	// r.Errno = RECODE_OK
	// r.Errmsg = RecodeErr(RECODE_OK)
	// resp["request_id"] = reqID
	defer r.RespData(&r.Resp)
	r.DelSession("nickname")
	r.DelSession("user_id")
}
