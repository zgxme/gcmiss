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
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", mysqluser, mysqlpass, mysqlurl, mysqlpor, mysqldb)
	_ = orm.RegisterDataBase("default", "mysql", dataSource, maxIdle, maxConn)
	beego.Info("connect db success!")
	orm.DefaultTimeLoc = time.UTC
	orm.RegisterModelWithPrefix(tablePre, new(User), new(Profile), new(Manager), new(Post))
	_ = orm.RunSyncdb("default", true, true)

}

//返回带前缀的表明
func TableNmae(str string) string {
	return beego.AppConfig.String("tb_") + str
}
