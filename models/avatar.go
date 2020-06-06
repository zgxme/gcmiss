/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2019-12-14 10:47:16
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-02 23:40:37
 */
package models

import "time"

//头像
type Avatar struct {
	Id         int64     //头像id
	AvatarName string    //图片名称
	Url        string    //图片url
	User       *User     `orm:"reverse(one)"`                //设置一对一的反向关系
	Ctime      time.Time `orm:"auto_now_add;type(datetime)"` //创建时间
	Mtime      time.Time `orm:"auto_now;type(datetime)"`     //修改时间
}
