/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2020-02-15 15:23:58
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime : 2020-02-15 15:41:29
 */
package models

import (
	"github.com/astaxie/beego/orm"
)

//图片
type Image struct {
	Id        int64  //图片id
	ImageName string //图片名称
	Url       string //图片url
	User      *User  `json:"user" orm:"rel(fk);index"` //设置一对多的反向关系
	Post      *Post  `json:"post" orm:"rel(fk);index"` //设置一对多的反向关系 帖子
	// Artical   *Artical `json:"artical" orm:"rel(fk);index"` //设置一对多的反向关系 物品
}

func AddImage(image *Image) error {
	o := orm.NewOrm()
	_, err := o.Insert(image)
	return err
}
