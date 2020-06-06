/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2020-02-15 15:23:58
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-23 17:24:36
 */
package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

//图片
type Image struct {
	Id        int64     //图片id
	ImageName string    //图片名称
	Url       string    //图片url
	User      *User     `json:"user" orm:"rel(fk);index"`    //设置一对多的反向关系
	Post      *Post     `json:"post" orm:"rel(fk);index"`    //设置一对多的反向关系 帖子
	Artical   *Artical  `json:"artical" orm:"rel(fk);index"` //设置一对多的反向关系 物品
	Ctime     time.Time `orm:"auto_now_add;type(datetime)"`  //创建时间
	Mtime     time.Time `orm:"auto_now;type(datetime)"`      //修改时间
}

type ImageItem struct {
	ImageID   int64  `json:"image_id"`
	ImageName string `json:image_name`
	ImageUrl  string `json:"image_url"`
}

func GetPostImageList(commentInfo map[string]interface{}) (error, *[]ImageItem) {
	o := orm.NewOrm()
	var sql string
	var ImageList []ImageItem
	postId := commentInfo["postId"].(int64)
	articalId := commentInfo["articalId"].(int64)
	if postId != 0 {
		sql = fmt.Sprintf("SELECT id,image_name,url FROM tb_image WHERE post_id = %d AND post_id IN (SELECT id FROM tb_post WHERE status != 1);", postId)
	} else {
		sql = fmt.Sprintf("SELECT id,image_name,url FROM tb_image WHERE artical_id = %d AND artical_id IN (SELECT id FROM tb_artical WHERE status != 1);", articalId)
	}
	var terms []orm.Params
	_, err := o.Raw(sql).Values(&terms)
	if err != nil {
		return err, &ImageList
	}
	for _, term := range terms {
		imageIDValue := term["id"].(string)
		imageID, _ := strconv.ParseInt(imageIDValue, 10, 64)
		imageUrl := term["url"].(string)
		imageName := term["image_name"].(string)
		image := ImageItem{
			ImageID:   imageID,
			ImageName: imageName,
			ImageUrl:  imageUrl,
		}
		ImageList = append(ImageList, image)
	}
	return err, &ImageList

}

func AddImage(image *Image) error {
	o := orm.NewOrm()
	_, err := o.Insert(image)
	return err
}
