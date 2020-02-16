/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2020-02-02 23:20:43
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime : 2020-02-09 11:15:13
 */
package controllers

import (
	"crypto/md5"
	"fmt"
	"gcmiss/models"
	. "gcmiss/models"
	"io"
	"math/rand"
	"os"
	"path"
	"time"

	"github.com/astaxie/beego"
)

//AddPostController add post controller
type AddPostController struct {
	BaseController
}

//AddPost add post
func (r *AddPostController) AddPost() {
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
	//TODO upload to qiniu
	var images []*Image
	if files, err := r.GetFiles("images"); err == nil {
		for i, _ := range files {
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				r.Errno = FILE_ERROR
				r.Errmsg = RecodeErr(FILE_ERROR)
				return
			}
			ext := path.Ext(files[i].Filename)
			if _, ok := AllowExtMap[ext]; !ok {
				r.Errno = FORMAT_ERROR
				r.Errmsg = RecodeErr(FORMAT_ERROR)
				beego.Error(r.Errmsg)
				return
			}
			//创建目录
			uploadDir := "/images/upload/" + time.Now().Format("2006/01/02/")
			if err := os.MkdirAll(uploadDir, 777); err != nil {
				r.Errno = FILE_ERROR
				r.Errmsg = RecodeErr(FILE_ERROR)
				beego.Error(r.Errmsg)
				return
			}
			//构造文件名称
			rand.Seed(time.Now().UnixNano())
			randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000)
			hashName := md5.Sum([]byte(time.Now().Format("2006_01_02_15_04_05_") + randNum))
			fileName := fmt.Sprintf("%x", hashName) + ext
			fpath := uploadDir + fileName
			//create destination file making sure the path is writeable.
			dst, err := os.Create(fpath)
			defer dst.Close()
			if err != nil {
				r.Errno = FILE_ERROR
				r.Errmsg = RecodeErr(FILE_ERROR)
				beego.Error(r.Errmsg)
				return
			}
			//copy the uploaded file to the destination file
			if _, err := io.Copy(dst, file); err != nil {
				r.Errno = FILE_ERROR
				r.Errmsg = RecodeErr(FILE_ERROR)
				beego.Error(r.Errmsg)
				return
			}
			images = append(images, &Image{ImageName: fileName, Url: fpath})
		}
	}
	nickName := r.GetSession("nickname")
	userID := r.GetSession("user_id")
	if userID == nil || nickName == nil {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(RecodeErr(PARAM_ERROR)))
		return
	}
	if err := models.AddPost(userID.(int64), &post, images); err != nil {
		r.Errno = DB_ERROR
		r.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
}
