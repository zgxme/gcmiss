package models

//学生身份
type Profile struct {
	Id         int64  ////beego默认Id为主键,且自增长
	Name       string //姓名
	StuId      string //学号
	School     string //学校
	Profession string //专业
	grade      string //年级
	Sex        int    //性别
	QQNumber   string //qq号
	Email      string //邮箱
	TelNumber  string //手机号
	User       *User  `orm:"reverse(one)"` //设置一对一的反向关系
}
