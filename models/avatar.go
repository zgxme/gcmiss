/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2019-12-14 10:47:16
 * @LastEditors  : Zheng Gaoxiong
 * @LastEditTime : 2020-02-15 22:10:32
 */
package models

//头像
type Avatar struct {
	Id         int64  //头像id
	AvatarName string //图片名称
	Url        string //图片url
	User       *User  `orm:"reverse(one)"` //设置一对一的反向关系

}
