/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2020-02-09 00:11:06
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-24 14:09:00
 */
package models

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type ArticalReq struct {
	Title   string `form:"artical_name"`
	Content string `form:"artical_desc"`
	Price   int    `form:"artical_price"`
}

type ArticalItem struct {
	ArticalID    int64  `json:"post_id"`
	Name         string `json:"title"`
	Description  string `json:"content"`
	Price        int    `json:"artical_price"`
	OwnerID      int64  `json:"poster_id"`
	OwnerName    string `json:"poster_name"`
	OwnerAvatar  string `json:"poster_avatar"`
	Ctime        string `json:"ctime"`
	ArticalLable string `json:"post_lable"`
}

//物品
type Artical struct {
	Id          int64  //beego默认Id为主键,且自增长
	Name        string //物品名称
	Description string //物品描述
	Price       int    //物品价格
	Status      int8   //0正常 1被删除
	User        *User  `json:"user" orm:"rel(fk);index"` //设置一对多的反向关系 物品创建人
	// Image       []*Image  `json:"images" orm:"reverse(many)"` //帖子的图片
	Ctime time.Time `orm:"auto_now_add;type(datetime)"` //创建时间
	Mtime time.Time `orm:"auto_now;type(datetime)"`     //修改时间
	Image []*Image  `json:"images" orm:"reverse(many)"` //帖子的图片
}

//发布物品
func AddArtical(userId int64, articalInfo *ArticalReq, images []*Image) error {
	o := orm.NewOrm()
	user, err := GetUser(userId)
	if err != nil {
		return err
	}
	// fmt.Println("price is", articalInfo.Price)
	var artical Artical
	artical.User = user
	artical.Name = articalInfo.Title
	artical.Description = articalInfo.Content
	artical.Price = articalInfo.Price
	artical.Ctime = time.Now()
	artical.Mtime = time.Now()
	artical.Status = 0
	_, err = o.Insert(&artical)

	var post Post
	for _, image := range images {
		image.User = user
		image.Artical = &artical
		image.Post = &post
		if err := AddImage(image); err != nil {
			return err
		}
	}
	return err
}

//删除物品
func DeleteArtical(userId int64, articalInfo map[string]interface{}) error {
	o := orm.NewOrm()
	articalID := int64(articalInfo["artical_id"].(float64))
	artical := Artical{Id: articalID}
	err := o.Read(&artical)
	if err != nil {
		return err
	}
	var exist bool
	exist, err = OneManagerExist(userId)
	// fmt.Println(exist, err)
	if err != nil {
		return err
	}

	if artical.User.Id == userId || exist {
		artical.Status = 1
		_, err = o.Update(&artical)
	}
	if !exist {
		return errors.New("not auth")
	}
	return err
}

//更新物品信息
func UpdateArtical(userId int64, articalInfo map[string]interface{}) error {
	o := orm.NewOrm()
	// fmt.Println("articalInfo", articalInfo)
	articalID := int64(articalInfo["artical_id"].(float64))
	articalName := articalInfo["artical_name"].(string)
	articalDesc := articalInfo["artical_desc"].(string)
	articalPrice := articalInfo["artical_price"].(string)
	price, _ := strconv.Atoi(articalPrice)
	artical := Artical{Id: articalID}
	err := o.Read(&artical)
	if err != nil {
		return err
	}
	if artical.User.Id == userId {
		artical.Name = articalName
		artical.Description = articalDesc
		artical.Price = price
		artical.Mtime = time.Now()
		_, err = o.Update(&artical)
	} else {
		return errors.New("not auth")
	}
	return err
}

//获取物品
func GetArtical(articalInfo map[string]interface{}) (error, int64, *[]ArticalItem) {
	o := orm.NewOrm()
	cursor := articalInfo["cursor"].(int64)
	limit := articalInfo["limit"].(int64)
	desc := articalInfo["desc"].(int)
	var sql string
	var articalList []ArticalItem
	sql = fmt.Sprintf("SELECT id,user_id,name,description,price,ctime FROM tb_artical WHERE status != 1 AND user_id IN (SELECT id FROM tb_user WHERE status = 0) ORDER BY mtime %s,id ASC LIMIT %d, %d", sqlDesc[desc], cursor, limit)
	var terms []orm.Params
	_, err := o.Raw(sql).Values(&terms)
	if err != nil {
		return err, 0, &articalList
	}
	for _, term := range terms {
		userIDValue := term["user_id"].(string)
		userID, _ := strconv.ParseInt(userIDValue, 10, 64)
		user := User{Id: userID}
		err = o.Read(&user)
		if err != nil {
			return err, 0, &articalList
		}
		priceValue, _ := term["price"].(string)
		// fmt.Println("sql", sql)
		// fmt.Println("priceValue", priceValue)
		price, _ := strconv.Atoi(priceValue)
		ownerName := user.Nickname
		articalIDValue := term["id"].(string)
		articalID, _ := strconv.ParseInt(articalIDValue, 10, 64)
		articalName := term["name"].(string)
		articaldescription := term["description"].(string)
		ctime := term["ctime"].(string)
		postLable := fmt.Sprintf("%s   %s   发表在[二手市场]", ownerName, ctime[:10])
		userInfo, _ := GetUserInfo(userID)
		artical := ArticalItem{
			ArticalID:    articalID,
			Name:         articalName,
			Description:  articaldescription,
			Price:        price,
			OwnerID:      userID,
			OwnerName:    ownerName,
			OwnerAvatar:  userInfo.Url,
			Ctime:        ctime,
			ArticalLable: postLable,
		}
		articalList = append(articalList, artical)
	}

	return err, int64(len(articalList)), &articalList
}

func GetArticalItem(articalId int64) (error, *ArticalItem) {
	o := orm.NewOrm()
	var ArticalSql string
	ArticalSql = fmt.Sprintf("SELECT id,user_id,name,description, price,ctime FROM tb_artical WHERE status != 1 AND id=%d AND user_id IN (SELECT id FROM tb_user WHERE status = 0);", articalId)
	var artical ArticalItem
	var terms []orm.Params
	_, err := o.Raw(ArticalSql).Values(&terms)
	if err != nil {
		return err, &artical
	}
	for _, term := range terms {
		userIDValue := term["user_id"].(string)
		userID, _ := strconv.ParseInt(userIDValue, 10, 64)
		user := User{Id: userID}
		err = o.Read(&user)
		if err != nil {
			return err, &artical
		}
		posterName := user.Nickname
		postIDValue := term["id"].(string)
		postID, _ := strconv.ParseInt(postIDValue, 10, 64)
		title := term["name"].(string)
		content := term["description"].(string)
		ctime := term["ctime"].(string)
		userInfo, _ := GetUserInfo(userID)
		// tagValue := term["tag"].(string)
		// tag, _ := strconv.Atoi(tagValue)
		postLable := fmt.Sprintf("%s   %s   发表在[二手市场]", posterName, ctime[:10])
		priceValue, _ := term["price"].(string)
		// fmt.Println("sql", sql)
		// fmt.Println("priceValue", priceValue)
		price, _ := strconv.Atoi(priceValue)
		artical = ArticalItem{
			ArticalID:    postID,
			Name:         title,
			Description:  content,
			Ctime:        ctime,
			OwnerID:      userID,
			OwnerName:    posterName,
			OwnerAvatar:  userInfo.Url,
			ArticalLable: postLable,
			Price:        price,
		}
		return err, &artical
	}
	return err, &artical
}
