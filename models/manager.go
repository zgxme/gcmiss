/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2019-12-14 10:49:43
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-24 11:40:04
 */
package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

//管理员表
type Manager struct {
	Id     int64 //beego默认Id为主键,且自增长
	Status int   //管理员权限
	//normal status = 0
	//manager status = 1
	//admin status = 2
	User  *User     `orm:"reverse(one)"`                //设置一对一的反向关系
	Ctime time.Time `orm:"auto_now_add;type(datetime)"` //创建时间
	Mtime time.Time `orm:"auto_now;type(datetime)"`     //修改时间
}

//check one manager exist or not
func OneManagerExist(UserId int64) (bool, error) {
	o := orm.NewOrm()
	var sql string
	//tb_user status is 0
	//tb_manager status is 1
	sql = fmt.Sprintf("SELECT id, nickname, status FROM tb_user WHERE status = 0 AND id = %d AND manager_id IN (SELECT id FROM tb_manager WHERE tb_manager.status != 0);", UserId)
	// fmt.Println(sql)
	var terms []orm.Params
	_, err := o.Raw(sql).Values(&terms)
	lenOfUser := len(terms)
	if err != nil {
		return false, err
	}
	if lenOfUser != 0 {
		return true, err
	}
	return false, err
}

//register manager
func RegisterManager(userId int64) error {
	o := orm.NewOrm()
	user := User{Id: userId}
	err := o.Read(&user)
	if err != nil {
		return err
	}
	manager := Manager{Id: user.Manager.Id}
	manager.Status = 1
	// fmt.Println(user)
	_, err = o.Update(&manager)
	return err
}

//delete manager
func DeleteManager(userId int64) error {
	o := orm.NewOrm()
	user := User{Id: userId}
	err := o.Read(&user)
	if err != nil {
		return err
	}
	manager := Manager{Id: user.Manager.Id}
	manager.Status = 0
	// fmt.Println(user)
	_, err = o.Update(&manager)
	return err
}
