/*
 * @Descripttion:add artical api
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2020-02-09 11:14:43
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-24 20:48:52
 */
package controllers

import (
	"crypto/md5"
	"fmt"
	"gcmiss/models"
	. "gcmiss/models"
	"math/rand"
	"time"

	"github.com/astaxie/beego"
)

func checkPrice(param int) bool {
	if param < 0 || param >= 1e9 {
		return false
	}
	return true
}

// @Title articalAdd
// @Description user add artical
// @Param artical_name	formData	string	true	"The name of artical"
// @Param	artical_desc	formData	string	true	"The desc of artical"
// @Param artical_price	formData	float false "The price of artical"
// @Param images	formData	file	true	"the images of post"
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure	2	param error
// @Failure 4002	db data error
// @Failure 4009	db error
// @Failure 4012	get file error
// @Failure	4013	file format error
// @router /add [post]
func (r *ArticalController) AddArtical() {
	// articalInfo := make(map[string]interface{})
	defer r.RespData(&r.Resp)
	// err := json.Unmarshal(r.Ctx.Input.RequestBody, &articalInfo)
	// if err != nil {
	// 	r.Errno = PARAM_ERROR
	// 	r.Errmsg = RecodeErr(PARAM_ERROR)
	// 	beego.Error(r.errLog(err.Error()))
	// 	return
	// }
	artical := ArticalReq{}
	//获取当前用户
	if err := r.ParseForm(&artical); err != nil {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
	//TODO check param
	if !checkTitle(artical.Title) {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(RecodeErr(PARAM_ERROR)))
		return
	}
	if !checkContent(artical.Content) {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(RecodeErr(PARAM_ERROR)))
		return
	}
	if !checkPrice(artical.Price) {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(RecodeErr(PARAM_ERROR)))
		return
	}

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
			fileUrl := fmt.Sprintf("https://%s.%s/%s/%s", r.buckname, r.endpoint, ARTICAL, fileName)
			fileSave := fmt.Sprintf("%s/%s", ARTICAL, fileName)
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
	// fmt.Println("artical", artical)
	err = models.AddArtical(userID.(int64), &artical, images)
	if err != nil {
		r.Errno = DB_ERROR
		r.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
}
