/**
 * @Author: zhenggaoxiong
 * @Description:
 * @File:  base
 * @Date: 2019/12/15 14:42
 */

package session

import "github.com/astaxie/beego"

func Init() {
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:6379"
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 86400

}
