/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2020-05-05 12:01:34
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-10 21:45:23
 */
package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
)

func registerSet(nickname string, password string, email string) (string, error) {
	redisHost := beego.AppConfig.String("redishost")
	redisPost, _ := beego.AppConfig.Int("redisport")
	expire := beego.AppConfig.String("registexist")
	redisURL := fmt.Sprintf("%s:%d", redisHost, redisPost)
	c, err := redis.Dial("tcp", redisURL)
	if err != nil {
		return "", err
	}
	defer c.Close()
	mockname := GetMd5String(nickname)
	timeStr := GetMd5String(time.Now().Format("2006-01-02 15:04:05"))
	key := mockname + timeStr
	// fmt.Println(key)
	password = GetPassword(password)
	value := fmt.Sprintf("%s %s %s", nickname, password, email)
	_, err = c.Do("SET", key, value, "EX", expire)
	if err != nil {
		return "", err
	}
	return key, err

	// if err != nil {
	// 	fmt.Println("redis get failed:", err)
	// } else {
	// 	fmt.Printf("Get mykey: %v \n", username)
	// }
}

func checkRegister(auth string) ([]string, error) {
	redisHost := beego.AppConfig.String("redishost")
	redisPost, _ := beego.AppConfig.Int("redisport")
	redisURL := fmt.Sprintf("%s:%d", redisHost, redisPost)
	c, err := redis.Dial("tcp", redisURL)
	defer c.Close()
	var infoAns []string
	if err != nil {
		return infoAns, err
	}
	info, err := redis.String(c.Do("GET", auth))
	if err != nil {
		return infoAns, err
	}
	if info != "" {
		// fmt.Println(info)
		infoAns = strings.Split(info, " ")
		// fmt.Println(infoAns)
	}
	if len(infoAns) == 3 {
		_, err = c.Do("DEL", auth)
	}
	return infoAns, err
}
