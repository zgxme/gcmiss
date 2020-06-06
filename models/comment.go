/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2020-04-04 13:29:05
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-23 19:22:18
 */

package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

//评论
type Comment struct {
	Id        int64  //评论id
	Content   string //评论内容
	PostId    int64  //帖子ID
	ArticalId int64  //物品ID
	UserFrom  *User  `json:"user" orm:"rel(fk);index"`
	UserTo    *User  `json:"user" orm:"rel(fk);index"`
	Status    int
	Ctime     time.Time `orm:"auto_now_add;type(datetime)"` //创建时间
	Mtime     time.Time `orm:"auto_now;type(datetime)"`     //修改时间
}

type CommentItem struct {
	CommentId      int64  `json:"comment_id"`
	Content        string `json:"content"`
	Ctime          string `json:"ctime"`
	UserFromId     int64  `json:"user_from_id"`
	UserFromName   string `json:"user_from_name"`
	UserFromAvatar string `json:"user_from_avatar"`
	UserToId       int64  `json:"user_to_id"`
	UserToName     string `json:"user_to_name"`
	UserToAvatar   string `json:"user_to_avatar"`
}

func AddComment(userId int64, userTo int64, content string, postId int64, articalId int64) error {
	o := orm.NewOrm()
	userFrom, err := GetUser(userId)
	if err != nil {
		return err
	}
	userTO, err := GetUser(userTo)
	if err != nil {
		return err
	}
	var comment Comment
	comment.UserFrom = userFrom
	comment.UserTo = userTO
	comment.Content = content
	comment.Ctime = time.Now()
	comment.Mtime = time.Now()
	comment.Status = 0
	comment.PostId = postId
	comment.ArticalId = articalId
	if _, err := o.Insert(&comment); err != nil {
		return err
	}
	return err
}

func GetCommentList(commentInfo map[string]interface{}) (error, int64, *[]CommentItem) {
	o := orm.NewOrm()
	var sql string
	var commentList []CommentItem
	postId := commentInfo["postId"].(int64)
	articalId := commentInfo["articalId"].(int64)
	cursor := commentInfo["cursor"].(int64)
	limit := commentInfo["limit"].(int64)
	desc := commentInfo["desc"].(int)
	if postId != 0 {
		sql = fmt.Sprintf("SELECT id,content,post_id,artical_id,user_from_id,user_to_id,ctime FROM tb_comment WHERE status != 1 AND post_id = %d AND post_id IN (SELECT id FROM tb_post WHERE status != 1) AND user_from_id IN (SELECT id FROM tb_user WHERE status != 1) AND user_to_id IN (SELECT id FROM tb_user WHERE status != 1) ORDER BY mtime %s,id ASC LIMIT %d, %d", postId, sqlDesc[desc], cursor, limit)
	} else {
		sql = fmt.Sprintf("SELECT id,content,post_id,artical_id,user_from_id,user_to_id,ctime FROM tb_comment WHERE status != 1 AND artical_id = %d AND artical_id IN (SELECT id FROM tb_artical WHERE status != 1) AND user_from_id IN (SELECT id FROM tb_user WHERE status != 1) AND user_to_id IN (SELECT id FROM tb_user WHERE status != 1) ORDER BY mtime %s,id ASC LIMIT %d, %d", articalId, sqlDesc[desc], cursor, limit)
	}
	var terms []orm.Params
	_, err := o.Raw(sql).Values(&terms)
	if err != nil {
		return err, int64(len(commentList)), &commentList
	}
	for _, term := range terms {
		userIDValue := term["user_from_id"].(string)
		userID, _ := strconv.ParseInt(userIDValue, 10, 64)
		userToIDValue := term["user_to_id"].(string)
		userToID, _ := strconv.ParseInt(userToIDValue, 10, 64)
		user := User{Id: userID}
		userTo := User{Id: userToID}
		err = o.Read(&user)
		err = o.Read(&userTo)
		if err != nil {
			return err, int64(len(commentList)), &commentList
		}
		userFromInfo, _ := GetUserInfo(userID)
		userToInfo, _ := GetUserInfo(userToID)
		userFromId := userID
		userFromName := user.Nickname
		userFromAvatar := userFromInfo.Url
		userToId := userToID
		userToName := userTo.Nickname
		userToAvatar := userToInfo.Url
		CommentIDValue := term["id"].(string)
		CommentID, _ := strconv.ParseInt(CommentIDValue, 10, 64)
		content := term["content"].(string)
		ctime := term["ctime"].(string)
		comment := CommentItem{
			CommentId:      CommentID,
			Content:        content,
			Ctime:          ctime,
			UserFromId:     userFromId,
			UserFromName:   userFromName,
			UserFromAvatar: userFromAvatar,
			UserToId:       userToId,
			UserToName:     userToName,
			UserToAvatar:   userToAvatar,
		}
		commentList = append(commentList, comment)
	}
	return err, int64(len(commentList)), &commentList
}
