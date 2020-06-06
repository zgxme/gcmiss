package models

/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2019-12-14 10:25:34
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-24 13:01:36
 */

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

const PROFILE_TABNAME = "tb_profile"

//用户
type User struct {
	Id       int64  //id increase
	Nickname string //nickname is not unique
	Password string //password
	Status   int64  //auth status
	//0 noraml
	//1 deleted
	//TODO status 2 is auth
	Avatar   *Avatar    `orm:"rel(one)"` //头像 设置一对一的反向关系
	Profile  *Profile   `orm:"rel(one)"` //真实身份 设置一对一的反向关系
	Manager  *Manager   `orm:"rel(one)"`
	Image    []*Image   `json:"images" orm:"reverse(many)"`
	Articals []*Artical `json:"articals" orm:"reverse(many)"`
	Posts    []*Post    `json:"posts" orm:"reverse(many)"`
	Comments []*Comment `json:"comments" orm:"reverse(many)"`
	Ctime    time.Time  `orm:"auto_now_add;type(datetime)"` //创建时间
	Mtime    time.Time  `orm:"auto_now;type(datetime)"`     //修改时间
}
type ProfileItem struct {
	ID         int64  `form:"-"`
	Name       string `form:"name"`
	StuID      string `form:"stu_id"`
	School     string `form:"school"`
	Profession string `form:"profession"`
	// grade      string `form:"grade"`
	Sex      int    `form:"sex"`
	QQNumber string `form:"qq_number"`
	// Email      string `form:"email"`
	TelNumber  string `form:"telnumber"`
	AvatarName string `form:"-"`
	Avatar     string `form:"-"`
	// Avatar     interface{} `form:"avatar"`
}

type UserItem struct {
	ID         int64  `json:"user_id"`
	Nickname   string `json:"nickname"`
	AvatarName string `json:"avatar_name"`
	Url        string `json:"avatar_url"`
	Status     int    `json:"manager_status"`
}

//check one user exist or not
func OneUserExist(nickName string) (bool, error) {
	o := orm.NewOrm()
	var sql string
	//status is not 1
	sql = fmt.Sprintf("SELECT id, nickname, status FROM tb_user WHERE status = 0 AND nickName = '%s';", nickName)
	// fmt.Println(sql)
	var terms []orm.Params
	_, err := o.Raw(sql).Values(&terms)
	lenOfUser := len(terms)

	if err != nil {
		return true, err
	}
	if lenOfUser != 0 {
		return true, err
	}
	return false, err
}

//check one email exist or not
func OneEmailExist(email string) (bool, error) {
	o := orm.NewOrm()
	var sql string
	sql = fmt.Sprintf("SELECT id, nickname, status FROM tb_user WHERE status != 1 AND profile_id in (SELECT id FROM tb_profile WHERE email = '%s')", email)
	var terms []orm.Params
	_, err := o.Raw(sql).Values(&terms)
	lenOfUser := len(terms)
	// fmt.Println(terms)
	if err != nil {
		return true, err
	}
	if lenOfUser != 0 {
		return true, err
	}
	return false, err
}

//返回一个用户的id
func GetOneuserID(nickName string) (int64, error) {
	o := orm.NewOrm()
	user := User{}
	sql := fmt.Sprintf("SELECT id, nickname FROM tb_user WHERE nickname = '%s' AND status = 0", nickName)
	// fmt.Println(sql)
	var terms []orm.Params
	_, err := o.Raw(sql).Values(&terms)
	for _, term := range terms {
		userIDValue := term["id"].(string)
		userID, _ := strconv.ParseInt(userIDValue, 10, 64)
		user = User{Id: userID}
		// fmt.Println("userid", userID)
		err = o.Read(&user)
	}
	return user.Id, err
}

func UpdateProfile(userID interface{}, profile *ProfileItem, avatar *Avatar) error {
	// name := userInfo["name"].(string)
	// stuId := userInfo["stu_id"].(string)
	// school := userInfo["school"].(string)
	// profession := userInfo["profession"].(string)
	// sex, _ := strconv.Atoi(userInfo["sex"].(string))
	// qqNumber := userInfo["qq_number"].(string)
	// email := userInfo["email"].(string)
	// telNum := userInfo["telNum"].(string)
	userinfo := User{Id: userID.(int64)}
	o := orm.NewOrm()
	err := o.Read(&userinfo)
	if err != nil {
		return err
	}
	profileinfo := Profile{Id: userinfo.Profile.Id}
	err = o.Read(&profileinfo)
	if err != nil {
		return err
	}
	var name, stuId, school, profession, qqNumber, telNum string
	// fmt.Println("profile", profile)
	if profile.Name == "undefined" {
		name = profileinfo.Name
	} else {
		name = profile.Name
	}
	if profile.StuID == "undefined" {
		stuId = profileinfo.StuId
	} else {
		stuId = profile.StuID
	}
	if profile.School == "undefined" {
		school = profileinfo.School
	} else {
		school = profile.School
	}
	if profile.Profession == "undefined" {
		profession = profileinfo.Profession
	} else {
		profession = profile.Profession
	}
	sex := profile.Sex
	if profile.QQNumber == "undefined" {
		qqNumber = profileinfo.QQNumber
	} else {
		qqNumber = profile.QQNumber
	}
	// email := profile.Email
	if profile.TelNumber == "undefined" {
		telNum = profileinfo.TelNumber
	} else {
		telNum = profile.TelNumber
	}
	var user User
	err = o.QueryTable("tb_user").Filter("id", userID.(int64)).One(&user)
	proId := user.Profile.Id
	_, err = o.Raw("UPDATE tb_profile SET name = ?,stu_id = ?,school = ?,profession=?, sex = ?, q_q_number = ?,tel_number = ?, mtime = ? WHERE id = ?", name, stuId, school, profession, sex, qqNumber, telNum, time.Now(), proId).Exec()
	// fmt.Println(ret)
	if avatar.Url != "" {
		_, err = o.QueryTable("tb_avatar").Filter("id", user.Avatar.Id).Update(orm.Params{
			"url":         avatar.Url,
			"avatar_name": avatar.AvatarName,
		})
	}
	if err != nil {
		return err
	}
	return nil
}

func GetUser(userId int64) (*User, error) {
	user := User{Id: userId}
	o := orm.NewOrm()
	err := o.Read(&user)
	return &user, err
}

func GetUserInfo(userId int64) (*UserItem, error) {
	var userItem UserItem
	o := orm.NewOrm()
	userItem.ID = userId
	err := o.Raw("SELECT nickname, avatar_name, url, tb_manager.status FROM tb_user INNER JOIN tb_avatar ON tb_user.avatar_id = tb_avatar.id INNER JOIN tb_manager ON tb_manager.id = tb_user.manager_id AND tb_user.id = ? AND tb_user.status != 1 ;", userId).QueryRow(&userItem)
	// userItem.Url = strings.Replace(userItem.Url, "\\u0026", "&", -1)
	// fmt.Println(userItem.Url)
	return &userItem, err
}

//delete user
func DeleteUser(userId int64) error {
	o := orm.NewOrm()
	user := User{Id: userId}
	err := o.Read(&user)
	if err != nil {
		return err
	}
	user.Status = 1
	_, err = o.Update(&user)
	return err
}

func RegisterUser(nickName string, passWord string, email string, isAdmin bool) (int64, error) {
	o := orm.NewOrm()
	user := new(User)
	user.Nickname = nickName
	user.Password = passWord
	profile := new(Profile)
	profile.Email = email
	avatar := new(Avatar)
	md5Name := GetMd5String(nickName)
	avatar.AvatarName = GetMd5String(nickName + md5Name)
	avatar.Url = fmt.Sprintf("https://cdn.v2ex.com/gravatar/%s?&d=retro", avatar.AvatarName)
	manager := new(Manager)
	if isAdmin {
		manager.Status = ADMIN
	}
	user.Avatar = avatar
	user.Profile = profile
	user.Manager = manager
	user.Status = 0
	var Id int64
	_, err := o.Insert(profile)
	if err != nil {
		return Id, err
	}
	_, err = o.Insert(avatar)
	if err != nil {
		return Id, err
	}
	_, err = o.Insert(manager)
	if err != nil {
		return Id, err
	}

	Id, err = o.Insert(user)
	if err != nil {
		return Id, err
	}
	return Id, err
}
