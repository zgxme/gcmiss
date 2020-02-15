/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2020-02-09 00:11:06
 * @LastEditors  : Zheng Gaoxiong
 * @LastEditTime : 2020-02-15 15:30:37
 */
package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type ArticalItem struct {
	ArticalID   int64   `json:"artical_id"`
	Name        string  `json:"artical_name"`
	Description string  `json:"artical_desc"`
	Price       float64 `json:"artical_price"`
	OwnerID     int64   `json:"owner_id"`
	OwnerName   string  `json:"owner_name"`
	Ctime       string  `json:"ctime"`
}

//物品
type Artical struct {
	Id          int64     //beego默认Id为主键,且自增长
	Name        string    //物品名称
	Description string    //物品描述
	Price       float64   //物品价格
	Status      int8      //0正常 1被删除
	User        *User     `json:"user" orm:"rel(fk);index"`   //设置一对多的反向关系 物品创建人
	Image       []*Image  `json:"images" orm:"reverse(many)"` //帖子的图片
	Ctime       time.Time //创建时间
	Mtime       time.Time //修改时间
}

//发布物品
func AddArtical(userId int64, articalInfo map[string]interface{}) error {
	o := orm.NewOrm()
	user, err := GetUser(userId)
	if err != nil {
		return err
	}
	var artical Artical
	artical.User = user
	artical.Name = articalInfo["artical_name"].(string)
	artical.Description = articalInfo["artical_desc"].(string)
	artical.Price = articalInfo["artical_price"].(float64)
	artical.Ctime = time.Now()
	artical.Mtime = time.Now()
	artical.Status = 0
	_, err = o.Insert(&artical)
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
	if artical.User.Id == userId {
		artical.Status = 1
		_, err = o.Update(&artical)
	}
	return err
}

//更新物品信息
func UpdateArtical(userId int64, articalInfo map[string]interface{}) error {
	o := orm.NewOrm()
	articalID := int64(articalInfo["artical_id"].(float64))
	articalName := articalInfo["artical_name"].(string)
	articalDesc := articalInfo["artical_desc"].(string)
	articalPrice := articalInfo["artical_price"].(float64)
	artical := Artical{Id: articalID}
	err := o.Read(&artical)
	if err != nil {
		return err
	}
	if artical.User.Id == userId {
		artical.Name = articalName
		artical.Description = articalDesc
		artical.Price = articalPrice
		artical.Mtime = time.Now()
		_, err = o.Update(&artical)
	}
	return err
}

//获取物品
func GetArtical(articalInfo map[string]interface{}) (error, *[]ArticalItem) {
	o := orm.NewOrm()
	cursor := articalInfo["cursor"].(int64)
	limit := articalInfo["limit"].(int64)
	desc := articalInfo["desc"].(int)
	var sql string
	var articalList []ArticalItem
	sql = fmt.Sprintf("SELECT id,user_id,name,description,price,ctime FROM tb_artical WHERE status != 1 ORDER BY mtime %s LIMIT %d, %d", sqlDesc[desc], cursor, limit)
	var terms []orm.Params
	_, err := o.Raw(sql).Values(&terms)
	if err != nil {
		return err, &articalList
	}
	for _, term := range terms {
		userIDValue := term["user_id"].(string)
		userID, _ := strconv.ParseInt(userIDValue, 10, 64)
		user := User{Id: userID}
		err = o.Read(&user)
		if err != nil {
			return err, &articalList
		}
		priceValue := term["price"].(string)
		price, _ := strconv.ParseFloat(priceValue, 64)
		ownerName := user.Nickname
		articalIDValue := term["id"].(string)
		articalID, _ := strconv.ParseInt(articalIDValue, 10, 64)
		articalName := term["name"].(string)
		articaldescription := term["description"].(string)
		ctime := term["ctime"].(string)
		artical := ArticalItem{
			ArticalID:   articalID,
			Name:        articalName,
			Description: articaldescription,
			Price:       price,
			OwnerID:     userID,
			OwnerName:   ownerName,
			Ctime:       ctime,
		}
		articalList = append(articalList, artical)
	}
	return err, &articalList
}
