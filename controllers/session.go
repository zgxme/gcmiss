package controllers

import (
	"gcmiss/models"
)

//SessionController session controller
type SessionController struct {
	BaseController
}

//GetSessionData get session
func (r *SessionController) GetSessionData() {
	resp := make(map[string]interface{})
	reqID := r.GetreqID()
	defer r.RespDataV2(resp)
	user := models.User{}
	resp["errno"] = DB_ERROR
	resp["errmsg"] = RecodeErr(DB_ERROR)
	resp["request_id"] = reqID
	nickname := r.GetSession("nickname")
	userID := r.GetSession("user_id")
	if nickname != nil {
		user.Nickname = nickname.(string)
		resp["errno"] = RECODE_OK
		resp["present_user"] = nickname
		resp["present_id"] = userID
		resp["errmsg"] = recodeText[RECODE_OK]
	}
}
