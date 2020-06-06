/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2019-12-14 19:18:14
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-24 14:51:11
 */
package controllers

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils"
)

//GetMd5String 生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//UniqueId 生成Guid字串
func UniqueId() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

//GetNewPassword DONE salt algorithm
func GetNewPassword(pasd string) string {
	h := md5.New()
	salt := beego.AppConfig.String("salt")
	pasd = pasd + salt
	h.Write([]byte(pasd))
	clipherStr := h.Sum(nil)
	newPass := fmt.Sprintf("%s", hex.EncodeToString(clipherStr))
	return newPass
}

//GetPassword get password
func GetPassword(pasd string) string {
	return GetNewPassword(GetNewPassword(pasd))
}

//UploadFile upload file
func UploadFile(fd []byte, file_name string) error {
	endpoint := beego.AppConfig.String("endpoint")
	ak := beego.AppConfig.String("ak")
	sk := beego.AppConfig.String("sk")
	buckname := beego.AppConfig.String("buckname")
	client, err := oss.New(endpoint, ak, sk)
	if err != nil {
		return err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(buckname)
	if err != nil {
		return err
	}

	// // 读取本地文件。
	// fd, err := os.Open("<yourLocalFile>")
	// if err != nil {
	// 	return err
	// }
	// defer fd.Close()

	// 上传文件流。
	// err = bucket.PutObject(file_name, fd)
	err = bucket.PutObject(file_name, bytes.NewReader([]byte(fd)))
	return err
}

//DONE check email format can copy
func checkEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func adminAuth(userName string) bool {
	adminStr := beego.AppConfig.String("adminList")
	adminList := strings.Split(adminStr, ",")
	for _, value := range adminList {
		if value == userName {
			return true
		}
	}
	return false
}

func sendMail(nickname string, emailto string, auth string) error {
	password := beego.AppConfig.String("password")
	username := beego.AppConfig.String("username")
	emailhost := beego.AppConfig.String("emailhost")
	emailport, _ := beego.AppConfig.Int("emailport")
	activeurl := beego.AppConfig.String("activeurl")
	config := fmt.Sprintf(`{"username":"%s","password":"%s","host":"%s","port":%d}`, username, password, emailhost, emailport)
	email := utils.NewEMail(config)
	email.To = []string{emailto}
	email.From = "ctrlcer@foxmail.com"
	email.Subject = "gcmiss帐户激活邮件"
	url := activeurl + auth
	email.Text = fmt.Sprintf("%s您好!\n感谢您在工程小秘书注册帐号!\n激活帐号需要点击下面的链接:%s", nickname, url)
	email.HTML = ""
	// email.AttachFile("1.jpg") // 附件
	// email.AttachFile("1.jpg", "1") // 内嵌资源
	err := email.Send()
	if err != nil {
		return err
	}
	return err
}
