package main

import (
	"github.com/astaxie/beego"
	_ "yuwenlanzi/routers"
	"yuwenlanzi/tools"
)

func main() {

	// 定时刷新access_token
	tools.Start()





	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
