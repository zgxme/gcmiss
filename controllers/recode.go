/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2019-12-14 16:14:06
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-03-21 14:34:21
 */
package controllers

const (
	RECODE_OK        = 0
	PARAM_ERROR      = 2
	RECODE_DATAERR   = 4002
	DB_EXIST         = 4003
	DB_NOT_EXIST     = 4004
	DB_ADD_SELF      = 4005
	DB_ERROR         = 4009
	FILE_ERROR       = 4012
	FORMAT_ERROR     = 4013
	RECODE_UNKNOWERR = 4000
	AUTH_LOGIN       = 5004
)

var recodeText = map[int]string{
	RECODE_OK:        "success",
	DB_ERROR:         "db error",
	DB_EXIST:         "db exist",
	DB_NOT_EXIST:     "db not exist",
	DB_ADD_SELF:      "cannot add self",
	PARAM_ERROR:      "param error",
	RECODE_UNKNOWERR: "unknow error",
	RECODE_DATAERR:   "db data error",
	AUTH_LOGIN:       "auth error",
	FILE_ERROR:       "get file error",
	FORMAT_ERROR:     "file format error",
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
}
