/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2019-12-14 16:14:06
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-10 19:46:23
 */
package models

const (
	COMMUNICATE = 0
	LOSE        = 1
	FIND        = 2
	HELP        = 3
)

var recodeTag = map[int]string{
	COMMUNICATE: "校园交流",
	LOSE:        "失物招领",
	FIND:        "寻物启事",
	HELP:        "求人办事",
}
