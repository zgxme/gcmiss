/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2019-12-15 21:42:55
 * @LastEditors  : Zheng Gaoxiong
 * @LastEditTime : 2020-02-16 00:25:34
 */
/**
 * @Author: zhenggaoxiong
 * @Description: update user info
 * @File:  updateUser
 * @Date: 2019/12/15 21:42
 */
package controllers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"gcmiss/models"
	"math/rand"
	"os"
	"path"
	"time"

	"github.com/astaxie/beego"
)

type UpdateUserController struct {
	SessionController
	models.User
}
type Profile struct {
	ID         int64  `form:"-"`
	Name       string `form:"name"`
	StuID      string `form:"stu_id"`
	School     string `form:"school"`
	Profession string `form:"profession"`
	grade      string `form:"grade"`
	Sex        int    `form:"sex"`
	QQNumber   string `form:"qq_number"`
	Email      string `form:"email"`
	TelNumber  string `form:"telnumber"`
	// Avatar     interface{} `form:"avatar"`
}

//TODO user nickname update
func (r *UpdateUserController) UpdateUser() {
	userInfo := make(map[string]interface{})
	profile := Profile{}
	if err := r.ParseForm(&profile); err != nil {
		r.Errno = FILE_ERROR
		r.Errmsg = RecodeErr(FILE_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
	f, h, _ := r.GetFile("avatar")
	ext := path.Ext(h.Filename)
	//验证后缀名是否符合要求
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
	defer f.Close()
	if err := r.SaveToFile("avatar", fpath); err != nil {
		r.Errno = FILE_ERROR
		r.Errmsg = RecodeErr(FILE_ERROR)
		return
	}
	// fmt.Println(avatar)
	defer r.RespData(&r.Resp)
	r.Errno = RECODE_OK
	r.Errmsg = RecodeErr(RECODE_OK)
	err := json.Unmarshal(r.Ctx.Input.RequestBody, &userInfo)
	if err != nil {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
	//获取当前用户
	nickName := r.GetSession("nickname")
	userID := r.GetSession("user_id")
	if userID == nil || nickName == nil {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(RecodeErr(PARAM_ERROR)))
		return
	}
	//newNickname := userInfo["new_nickname"].(string)
	//exist := models.OneUserExist(newNickname)
	//if exist {
	//	r.Errno = DB_EXIST
	//	r.Errmsg = RecodeErr(DB_EXIST)
	//	beego.Error(u.errLog(RecodeErr(DB_EXIST)))
	//	return
	//}

	//if newNickname == ""{
	//	r.Errno = PARAM_ERROR
	//	r.Errmsg = RecodeErr(PARAM_ERROR)
	//	beego.Error(u.errLog(RecodeErr(PARAM_ERROR)))
	//	return
	//}
	err = models.UpdateProfile(userID, userInfo)
	if err != nil {
		r.Errno = DB_ERROR
		r.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
}
