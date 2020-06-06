package models

/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2019-12-14 10:46:44
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-23 14:22:52
 */

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type PostReq struct {
	Title   string `form:"title"`
	Content string `form:"content"`
	Tag     int    `form:"tag"`
}

type PostItem struct {
	PostId       int64  `json:"post_id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	Ctime        string `json:"ctime"`
	PosterId     int64  `json:"poster_id"`
	PosterName   string `json:"poster_name"`
	PosterAvatar string `json:"poster_avatar"`
	PostLable    string `json:"post_lable"`
}

//帖子
type Post struct {
	Id      int64     //beego默认Id为主键,且自增长
	Title   string    //标题
	Ctime   time.Time `orm:"auto_now_add;type(datetime)"` //创建时间
	Mtime   time.Time `orm:"auto_now;type(datetime)"`     //修改时间
	Content string    //帖子内容
	Status  int8      //0正常 1被删除
	Tag     int
	//0交流
	//1失物
	//2寻物
	//3寻求帮助
	User  *User    `json:"user" orm:"rel(fk);index"`   //设置一对多的反向关系 发帖人
	Image []*Image `json:"images" orm:"reverse(many)"` //帖子的图片
}

//发帖子
//TODO可以修改更加优雅的
func AddPost(userId int64, postInfo *PostReq, images []*Image) error {
	o := orm.NewOrm()
	user, err := GetUser(userId)
	if err != nil {
		return err
	}
	var post Post
	post.User = user
	post.Title = postInfo.Title
	post.Content = postInfo.Content
	post.Tag = postInfo.Tag
	post.Ctime = time.Now()
	post.Mtime = time.Now()
	post.Status = 0

	if _, err := o.Insert(&post); err != nil {
		return err
	}
	var artical Artical
	for _, image := range images {
		image.User = user
		image.Post = &post
		image.Artical = &artical
		if err := AddImage(image); err != nil {
			return err
		}
	}
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
	var exist bool
	exist, err = OneManagerExist(userId)
	// fmt.Println(exist, err)
	if err != nil {
		return err
	}

	// fmt.Println(post.User.Id)
	// fmt.Println(userId)
	// fmt.Println()
	if post.User.Id == userId || exist {
		post.Status = 1
		_, err = o.Update(&post)
		return err
	}
	if !exist {
		return errors.New("not auth")
	}
	return err
}

//更新帖子
//TODO:pass by pointer
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
	} else {
		return errors.New("not auth")
	}
	return err
}

//获取帖子
func GetPost(postInfo map[string]interface{}) (error, int64, *[]PostItem) {
	o := orm.NewOrm()
	cursor := postInfo["cursor"].(int64)
	limit := postInfo["limit"].(int64)
	desc := postInfo["desc"].(int)
	tag := postInfo["tag"].(int)
	var sql string
	var postList []PostItem
	sql = fmt.Sprintf("SELECT id,user_id,title,content,ctime, tag FROM tb_post WHERE status != 1 AND tag = %d AND user_id IN (SELECT id FROM tb_user WHERE status = 0) ORDER BY mtime %s,id ASC LIMIT %d, %d;", tag, sqlDesc[desc], cursor, limit)
	if tag == 5 {
		sql = fmt.Sprintf("SELECT id,user_id,title,content,ctime, tag FROM tb_post WHERE status != 1 AND user_id IN (SELECT id FROM tb_user WHERE status = 0) ORDER BY mtime %s,id ASC LIMIT %d, %d;", sqlDesc[desc], cursor, limit)
	}
	var terms []orm.Params
	_, err := o.Raw(sql).Values(&terms)
	if err != nil {
		return err, int64(len(postList)), &postList
	}
	for _, term := range terms {
		userIDValue := term["user_id"].(string)
		userID, _ := strconv.ParseInt(userIDValue, 10, 64)
		user := User{Id: userID}
		err = o.Read(&user)
		if err != nil {
			return err, int64(len(postList)), &postList
		}
		posterName := user.Nickname
		postIDValue := term["id"].(string)
		postID, _ := strconv.ParseInt(postIDValue, 10, 64)
		title := term["title"].(string)
		content := term["content"].(string)
		ctime := term["ctime"].(string)
		userInfo, _ := GetUserInfo(userID)
		tagValue := term["tag"].(string)
		tag, _ := strconv.Atoi(tagValue)
		postLable := fmt.Sprintf("%s   %s   发表在[%s]", posterName, ctime[:10], recodeTag[tag])
		post := PostItem{
			PostId:       postID,
			Title:        title,
			Content:      content,
			Ctime:        ctime,
			PosterId:     userID,
			PosterName:   posterName,
			PosterAvatar: userInfo.Url,
			PostLable:    postLable,
		}
		postList = append(postList, post)
	}
	return err, int64(len(postList)), &postList
}

func GetPostItem(postId int64) (error, *PostItem) {
	o := orm.NewOrm()
	var PostSql string
	PostSql = fmt.Sprintf("SELECT id,user_id,title,content,ctime, tag FROM tb_post WHERE status != 1 AND id=%d AND user_id IN (SELECT id FROM tb_user WHERE status = 0);", postId)
	var post PostItem
	var terms []orm.Params
	_, err := o.Raw(PostSql).Values(&terms)
	if err != nil {
		return err, &post
	}
	for _, term := range terms {
		userIDValue := term["user_id"].(string)
		userID, _ := strconv.ParseInt(userIDValue, 10, 64)
		user := User{Id: userID}
		err = o.Read(&user)
		if err != nil {
			return err, &post
		}
		posterName := user.Nickname
		postIDValue := term["id"].(string)
		postID, _ := strconv.ParseInt(postIDValue, 10, 64)
		title := term["title"].(string)
		content := term["content"].(string)
		ctime := term["ctime"].(string)
		userInfo, _ := GetUserInfo(userID)
		tagValue := term["tag"].(string)
		tag, _ := strconv.Atoi(tagValue)
		postLable := fmt.Sprintf("%s   %s   发表在[%s]", posterName, ctime[:10], recodeTag[tag])
		post = PostItem{
			PostId:       postID,
			Title:        title,
			Content:      content,
			Ctime:        ctime,
			PosterId:     userID,
			PosterName:   posterName,
			PosterAvatar: userInfo.Url,
			PostLable:    postLable,
		}
		return err, &post
	}
	return err, &post
}
