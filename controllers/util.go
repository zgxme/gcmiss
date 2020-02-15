/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2019-12-14 19:18:14
 * @LastEditors  : Zheng Gaoxiong
 * @LastEditTime : 2020-02-15 23:19:48
 */
package controllers

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/astaxie/beego"
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
