/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2020-02-08 15:12:37
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-23 23:25:03
 */

package models

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils"
)

const (
	DESC = 1
	ASC  = 0
)

var sqlDesc = map[int]string{
	DESC: "DESC",
	ASC:  "ASC",
}

const (
	NORMAL  = 0
	MANAGER = 1
	ADMIN   = 2
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

func sendMail(nickname string, emailto string, auth string) error {
	config := `{"username":"ctrlcer@foxmail.com","password":"ympevnmadaltbigb","host":"smtp.qq.com","port":587}`
	email := utils.NewEMail(config)
	email.To = []string{emailto}
	email.From = "ctrlcer@foxmail.com"
	email.Subject = "gcmiss帐户激活邮件"
	email.Text = fmt.Sprintf("%s您好，感谢您在工程小秘书注册帐户！激活帐户需要点击下面的链接:%s", nickname, auth)
	email.HTML = ""
	// email.AttachFile("1.jpg") // 附件
	// email.AttachFile("1.jpg", "1") // 内嵌资源
	err := email.Send()
	if err != nil {
		return err
	}
	return err
}
