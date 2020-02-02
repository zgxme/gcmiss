package models

//帖子
type Post struct {
	Id         int64  ////beego默认Id为主键,且自增长
	Title      string //
	StuId      string //学号
	School     string //学校
	Profession string //专业
	grade      string //年级
	Sex        int    //性别
	QQNumber   string //qq号
	Email      string //邮箱
	TelNumber  string //手机号
	}
