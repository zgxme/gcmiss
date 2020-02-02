package logs

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func Init() {
	//log 的配置
	//log打印文件名和行数
	beego.SetLogFuncCall(true)
	_ = beego.SetLogger(logs.AdapterMultiFile, `{"filename":"./gcmiss_logs/access.log","separate":["error", "info", "debug"],"maxdays":60,"color":true}`)
}
