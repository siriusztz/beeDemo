package main

import (
	_ "webtest1/routers"
	"github.com/astaxie/beego"
)


func main() {
	beego.SetStaticPath("/down1", "download1")

	//打开sessio
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime=10	//5秒时间
	beego.BConfig.WebConfig.Session.SessionProvider="file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"	//设置文件夹

	beego.Run()
}