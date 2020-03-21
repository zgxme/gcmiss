/**
 * @Author: zhenggaoxiong
 * @Description: logout api
 * @File:  logout
 * @Date: 2019/12/15 15:14
 */

package controllers

// @Title userLogout
// @Description user logout
// @Success 200 {object} models.ZDTCustomer.Customer
// @router /logout [post]
func (r *SessionController) Logout() {
	defer r.RespData(&r.Resp)
	r.DelSession("nickname")
	r.DelSession("user_id")
}
