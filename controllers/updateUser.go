/*
 * @Descripttion: user update api
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2019-12-15 21:42:55
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-24 12:40:28
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

func checkQQ(param string) bool {
	//check len
	if utf8.RuneCountInString(param) == 0 || param == "undefined" {
		return true
	}
	if utf8.RuneCountInString(param) > 15 {
		return false
	}
	//check space char not exist
	index := strings.Index(param, " ")
	if index != -1 {
		return false
	}
	return true
}

func checkTel(param string) bool {
	//check len
	if utf8.RuneCountInString(param) == 0 || param == "undefined" {
		return true
	}
	if utf8.RuneCountInString(param) != 11 {
		return false
	}
	//check space char not exist
	index := strings.Index(param, " ")
	if index != -1 {
		return false
	}
	return true
}

func checkName(param string) bool {
	//check len
	if utf8.RuneCountInString(param) == 0 || param == "undefined" {
		return true
	}
	if utf8.RuneCountInString(param) > 10 {
		return false
	}
	//check space char not exist
	index := strings.Index(param, " ")
	if index != -1 {
		return false
	}
	return true
}

func checkStuNum(param string) bool {
	//check len
	if utf8.RuneCountInString(param) == 0 || param == "undefined" {
		return true
	}
	if utf8.RuneCountInString(param) != 8 {
		return false
	}
	//check space char not exist
	index := strings.Index(param, " ")
	if index != -1 {
		return false
	}
	return true
}

func checkSexTag(param int) bool {
	if param < 0 || param > 1 {
		return false
	}
	return true
}

// @Title userRealInfoUpdate
// @Description user update real info
// @Param name	formData	string	false	"The real name of user"
// @Param stu_id	formData string false "The stu_id of user"
// @Param school	formData string false "The school of user"
// @Param profession	formData string false "The profession of user"
// @Param grade	formData string false "The grade of user"
// @Param sex formData int false "The sex of user"
// @Param qq_number	formData int false "The qq_number of user"
// @Param	email	formData string false "The email of user"
// @Param telnumber	formData string false "The telnumber of user"
// @Param	avatar	formData	file	"The avatar	of user"
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure	2	param error
// @Failure 4002	db data error
// @Failure 4009	db error
// @Failure 4012	get file error
// @Failure	4013	file format error
// @router /update [post]
func (r *UserController) UpdateUser() {
	defer r.RespData(&r.Resp)
	// userInfo := make(map[string]interface{})
	profile := ProfileItem{}
	if err := r.ParseForm(&profile); err != nil {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(err.Error()))
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
	if !checkQQ(profile.QQNumber) {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(RecodeErr(PARAM_ERROR)))
		return
	}
	if !checkTel(profile.TelNumber) {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(RecodeErr(PARAM_ERROR)))
		return
	}
	if !checkStuNum(profile.StuID) {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(RecodeErr(PARAM_ERROR)))
		return
	}
	if !checkName(profile.Name) {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(RecodeErr(PARAM_ERROR)))
		return
	}
	if !checkName(profile.School) {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(RecodeErr(PARAM_ERROR)))
		return
	}
	if !checkName(profile.Profession) {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(RecodeErr(PARAM_ERROR)))
		return
	}
	if !checkSexTag(profile.Sex) {
		r.Errno = PARAM_ERROR
		r.Errmsg = RecodeErr(PARAM_ERROR)
		beego.Error(r.errLog(RecodeErr(PARAM_ERROR)))
		return
	}
	var avatar Avatar
	file, header, err := r.GetFile("avatar")
	// fmt.Println(file)
	if err != nil {
		beego.Info(r.errLog(err.Error()))
	}
	if file != nil {
		file_op, err := header.Open()
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
		// uploadDir := "/avatar/upload/" + time.Now().Format("2006/01/02/")
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
		fileUrl := fmt.Sprintf("https://%s.%s/%s/%s", r.buckname, r.endpoint, AVATAR, fileName)
		fileSave := fmt.Sprintf("%s/%s", AVATAR, fileName)
		imageItem := make([]byte, 3145728)
		n, err := file_op.Read(imageItem)
		if err != nil {
			r.Errno = FILE_ERROR
			r.Errmsg = RecodeErr(FILE_ERROR)
			beego.Error(r.Errmsg)
			return
		}
		if !checkFileSize(n) {
			r.Errno = PARAM_ERROR
			r.Errmsg = RecodeErr(PARAM_ERROR)
			beego.Error(r.errLog(err.Error()))
			return
		}
		if n > 0 {
			imageItem = imageItem[:n]
		}
		// fmt.Println(imageItem)
		if err := UploadFile(imageItem, fileSave); err != nil {
			r.Errno = UPLOAD_OSS_ERROR
			r.Errmsg = RecodeErr(UPLOAD_OSS_ERROR)
			beego.Error(err.Error())
			return
		}

		avatar = Avatar{AvatarName: fileName, Url: fileUrl}
	}
	if err := models.UpdateProfile(userID, &profile, &avatar); err != nil {
		r.Errno = DB_ERROR
		r.Errmsg = RecodeErr(DB_ERROR)
		beego.Error(r.errLog(err.Error()))
		return
	}
}
