/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2019-12-14 10:37:58
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-23 13:24:52
 */

package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//real student info
type Profile struct {
	Id         int64     ////beego默认Id为主键,且自增长
	Name       string    //姓名
	StuId      string    //学号
	School     string    //学校
	Profession string    //专业
	grade      string    //年级
	Sex        int       //性别
	QQNumber   string    //qq号
	Email      string    //邮箱
	TelNumber  string    //手机号
	User       *User     `orm:"reverse(one)"`                //设置一对一的反向关系
	Ctime      time.Time `orm:"auto_now_add;type(datetime)"` //创建时间
	Mtime      time.Time `orm:"auto_now;type(datetime)"`     //修改时间
}

type ProfileItemResp struct {
	ID        int64  `json:"profile_id"`
	QQNumber  string `json:"qq_number"`
	TelNumber string `json:"telphone_num"`
	Email     string `json:"email"`
}

func GetProfileInfo(userId int64) (*ProfileItemResp, error) {
	var profileItem ProfileItemResp
	o := orm.NewOrm()
	profileItem.ID = userId
	err := o.Raw("SELECT q_q_number, tel_number, email FROM tb_profile WHERE id IN (SELECT profile_id FROM tb_user WHERE id = ? AND status = 0);", userId).QueryRow(&profileItem)
	// userItem.Url = strings.Replace(userItem.Url, "\\u0026", "&", -1)
	// fmt.Println(userItem.Url)
	return &profileItem, err
}
