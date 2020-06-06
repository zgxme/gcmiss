/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2019-12-14 16:14:06
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-23 14:12:44
 */
package controllers

const (
	AVATAR  = "avatar"
	POST    = "post"
	ARTICAL = "artical"
)

const (
	RECODE_OK        = 0
	PARAM_ERROR      = 2
	RECODE_DATAERR   = 4002
	DB_EXIST         = 4003
	DB_NOT_EXIST     = 4004
	DB_ADD_SELF      = 4005
	DB_DELETE_SELF   = 4006
	EMAIL_EXIST      = 4007
	DB_ERROR         = 4009
	REDIS_ERROR      = 4010
	SEND_EMAIL_ERROR = 4011
	FILE_ERROR       = 4012
	FORMAT_ERROR     = 4013
	UPLOAD_OSS_ERROR = 4014
	RECODE_UNKNOWERR = 4000
	AUTH_LOGIN       = 5004
	MOCK_TEST        = 1024
)

var recodeText = map[int]string{
	RECODE_OK:        "success",
	DB_ERROR:         "db error",
	DB_EXIST:         "db exist",
	DB_NOT_EXIST:     "db not exist",
	DB_ADD_SELF:      "cannot add self",
	DB_DELETE_SELF:   "cannot delete self",
	EMAIL_EXIST:      "email exist",
	PARAM_ERROR:      "param error",
	RECODE_UNKNOWERR: "unknow error",
	SEND_EMAIL_ERROR: "send email error",
	RECODE_DATAERR:   "db data error",
	REDIS_ERROR:      "redis error",
	AUTH_LOGIN:       "auth error",
	FILE_ERROR:       "get file error",
	FORMAT_ERROR:     "file format error",
	UPLOAD_OSS_ERROR: "upload oss error",
}

//RecodeErr get errmsg
func RecodeErr(code int) string {
	str, ok := recodeText[code]
	if ok {
		return str
	}
	return RecodeErr(RECODE_UNKNOWERR)
}

var AllowExtMap map[string]bool = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".JPG":  true,
	".JPEG": true,
	".PNG":  true,
}
