/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2019-12-14 10:46:44
 * @LastEditors  : Zheng Gaoxiong
 * @LastEditTime : 2020-02-08 02:37:27
 */

package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//帖子
type Post struct {
	Id      int64     //beego默认Id为主键,且自增长
	Title   string    //标题
	Ctime   time.Time //创建时间
	Mtime   time.Time //修改时间
	Content string    //帖子内容
	Status  int8      //0正常 1被删除
	User    *User     `json:"user" orm:"rel(fk);index"` //设置一对多的反向关系 发帖人
}

//发帖子
//TODO可以修改更加优雅的
func AddPost(userId int64, postInfo map[string]interface{}) error {
	o := orm.NewOrm()
	user, err := GetUser(userId)
	if err != nil {
		return err
	}
	var post Post
	post.User = user
	post.Title = postInfo["title"].(string)
	post.Content = postInfo["content"].(string)
	post.Ctime = time.Now()
	post.Mtime = time.Now()
	post.Status = 0
	_, err = o.Insert(&post)
	return err
}

//删除帖子
func DeletePost(userId int64, postInfo map[string]interface{}) error {
	o := orm.NewOrm()
	postID := int64(postInfo["post_id"].(float64))
	post := Post{Id: postID}
	err := o.Read(&post)
	if err != nil {
		return err
	}
	if post.User.Id == userId {
		post.Status = 1
		_, err = o.Update(&post)
	}
	return err
}

//更新帖子
func UpdatePost(userId int64, postInfo map[string]interface{}) error {
	o := orm.NewOrm()
	postID := int64(postInfo["post_id"].(float64))
	postContent := postInfo["content"].(string)
	postTitle := postInfo["title"].(string)
	post := Post{Id: postID}
	err := o.Read(&post)
	if err != nil {
		return err
	}
	if post.User.Id == userId {
		post.Content = postContent
		post.Title = postTitle
		post.Mtime = time.Now()
		_, err = o.Update(&post)
	}
	return err
}
