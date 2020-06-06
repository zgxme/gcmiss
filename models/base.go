/*
 * @Descripttion:
 * @version:
 * @Author: Zheng Gaoxiong
 * @Date: 2019-12-14 10:27:00
 * @LastEditors: Zheng Gaoxiong
 * @LastEditTime: 2020-05-02 11:22:06
 */
package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {

	mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpass := beego.AppConfig.String("mysqlpass")
	mysqlurl := beego.AppConfig.String("mysqlurl")
	mysqlpor, _ := beego.AppConfig.Int("mysqlport")
	mysqldb := beego.AppConfig.String("mysqldb")
	maxConn, _ := beego.AppConfig.Int("maxConn")
	maxIdle, _ := beego.AppConfig.Int("maxIdle")
	tablePre := beego.AppConfig.String("tablepre")
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=Asia%%2FShanghai", mysqluser, mysqlpass, mysqlurl, mysqlpor, mysqldb)
	_ = orm.RegisterDataBase("default", "mysql", dataSource, maxIdle, maxConn)
	beego.Info("connect db success!")
	orm.DefaultTimeLoc = time.UTC
	orm.RegisterModelWithPrefix(tablePre, new(User), new(Profile), new(Manager), new(Post), new(Artical), new(Avatar), new(Image), new(Comment))
	mysqlforce, _ := beego.AppConfig.Bool("mysqlforce")
	mysqlverbose, _ := beego.AppConfig.Bool("mysqlverbose")
	_ = orm.RunSyncdb("default", mysqlforce, mysqlverbose)

}

//返回带前缀的表明
func TableNmae(str string) string {
	return beego.AppConfig.String("tb_") + str
}
