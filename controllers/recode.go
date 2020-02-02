package controllers

const (
	RECODE_OK        = 0
	PARAM_ERROR      = 2
	RECODE_DATAERR   = 4002
	DB_EXIST         = 4003
	DB_NOT_EXIST     = 4004
	DB_ADD_SELF      = 4005
	DB_ERROR         = 4009
	RECODE_UNKNOWERR = 4000
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
}

//RecodeErr get errmsg
func RecodeErr(code int) string {
	str, ok := recodeText[code]
	if ok {
		return str
	}
	return RecodeErr(RECODE_UNKNOWERR)
}
