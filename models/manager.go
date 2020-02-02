package models

import "github.com/astaxie/beego/orm"

//管理员表
type Manager struct {
	Id     int64 //beego默认Id为主键,且自增长
	UserId int64 //用户id
	Status int   //管理员权限
	//User   *User `orm:"reverse(one)"` //设置一对一的反向关系
}

//判断一个管理员是否存在
func OneManagerExist(UserId int64) bool {
	o := orm.NewOrm()
	manager := Manager{}
	err := o.QueryTable("tb_manager").Filter("user_id", UserId).One(&manager)
	if err != orm.ErrNoRows {
		return true
	}
	return false
}
