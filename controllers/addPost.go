/*
 * @Descripttion: add post api
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2020-02-02 23:20:43
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-23 14:14:10
 */
package controllers

import (
	"crypto/md5"
	"fmt"
	"gcmiss/models"
	. "gcmiss/models"
	"math/rand"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/astaxie/beego"
)

//if true,check right
func checkFileSize(param int) bool {
	//check size
	if param >= 215040 {
		return false
	}
	return true
}

//if true,check right
func checkTitle(param string) bool {
	//check len
	if utf8.RuneCountInString(param) > 20 || utf8.RuneCountInString(param) < 5 {
		return false
	}
	//check space char not exist
	index := strings.Index(param, " ")
	if index != -1 {
		return false
	}
	return true
}

//if true,check right
func checkContent(param string) bool {
	//check len
	if utf8.RuneCountInString(param) > 200 || utf8.RuneCountInString(param) < 5 {
		return false
	}
	return true
}

func checkTag(param int) bool {
	if param < 0 || param > 4 {
		return false
	}
	return true
}

// @Title postAdd
// @Description user add post
// @Param name	formData	string	false	"The real name of user"
// @Param	title	formData	string	true	"The	title of post"
// @Param content formData	string	true	"The content of post"
// @Param tag fromData int true "The tag of post"
// @Param images	formData	file	true	"the images of post"
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure	2	param error
// @Failure 4002	db data error
// @Failure 4009	db error
// @Failure 4012	get file error
// @Failure	4013	file format error
// @router /add [post]
func (r *PostController) AddPost() {
	// postInfo := make(map[string]interface{})
	defer r.RespData(&r.Resp)
	// err := json.Unmarshal(r.Ctx.Input.RequestBody, &postInfo)
	// if err != nil {
	// 	r.Errno = PARAM_ERROR
	// 	r.Errmsg = RecodeErr(PARAM_ERROR)
	// 	beego.Error(r.errLog(err.Error()))
	// 	return
	// }
	//获取当前用户
	post := PostReq{}
	if err := r.ParseForm(&post); err != nil {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
	if !checkTitle(post.Title) {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(RecodeErr(PARAM_ERROR)))
		return
	}
	if !checkContent(post.Content) {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(RecodeErr(PARAM_ERROR)))
		return
	}
	if !checkTag(post.Tag) {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(RecodeErr(PARAM_ERROR)))
		return
	}

	//DONE upload to ali
	nickName := r.GetSession("nickname")
	userID := r.GetSession("user_id")
	if userID == nil || nickName == nil {
		r.Errno = AUTH_LOGIN
		r.Errmsg = RecodeErr(AUTH_LOGIN)
		beego.Error(r.errLog(RecodeErr(AUTH_LOGIN)))
		return
	}
	var images []*Image
	files, err := r.GetFiles("images")
	if err != nil {
		// r.Errno = PARAM_ERROR
		// r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Info(r.errLog(err.Error()))
	}
	if len(files) != 0 {
		for i, _ := range files {
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				r.Errno = FILE_ERROR
				r.Errmsg = RecodeErr(FILE_ERROR)
				return
			}
			// ext := path.Ext(files[i].Filename)
			// if _, ok := AllowExtMap[ext]; !ok {
			// 	r.Errno = FORMAT_ERROR
			// 	r.Errmsg = RecodeErr(FORMAT_ERROR)
			// 	beego.Error(r.Errmsg)
			// 	// return
			// }
			//创建目录
			// uploadDir := "/images/upload/" + time.Now().Format("2006/01/02/")
			// if err := os.MkdirAll(uploadDir, 777); err != nil {
			// 	r.Errno = FILE_ERROR
			// 	r.Errmsg = RecodeErr(FILE_ERROR)
			// 	beego.Error(r.Errmsg)
			// 	return
			// }
			rand.Seed(time.Now().UnixNano())
			randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000)
			hashName := md5.Sum([]byte(time.Now().Format("2006_01_02_15_04_05_") + randNum))
			preFileName := time.Now().Format("2006_01_02_15_04_05_")
			fileName := fmt.Sprintf("%s%x", preFileName, hashName) + ".JPG"
			fileUrl := fmt.Sprintf("https://%s.%s/%s/%s", r.buckname, r.endpoint, POST, fileName)
			fileSave := fmt.Sprintf("%s/%s", POST, fileName)
			imageItem := make([]byte, 3145728)
			n, err := file.Read(imageItem)
			if err != nil {
				r.Errno = FILE_ERROR
				r.Errmsg = RecodeErr(FILE_ERROR)
				beego.Error(r.Errmsg)
				return
			}
			if n > 0 {
				imageItem = imageItem[:n]
			}
			if !checkFileSize(n) {
				r.Errno = PARAM_ERROR
				r.Errmsg = RecodeErr(PARAM_ERROR)
				beego.Error(RecodeErr(PARAM_ERROR))
				return
			}
			if err := UploadFile(imageItem, fileSave); err != nil {
				r.Errno = UPLOAD_OSS_ERROR
				r.Errmsg = RecodeErr(UPLOAD_OSS_ERROR)
				beego.Error(err.Error())
				return
			}
			images = append(images, &Image{ImageName: fileName, Url: fileUrl})
		}
	}

	if err := models.AddPost(userID.(int64), &post, images); err != nil {
		r.Errno = DB_ERROR
		r.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
}
