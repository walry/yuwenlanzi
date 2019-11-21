package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "yuwenlanzi/routers"
	"yuwenlanzi/service/wechat"
)

func main() {
	logPath := fmt.Sprintf("{\"filename\":\"%s\"}",beego.AppConfig.String("logpath") + "/yuwenlanzi.log")
	err := logs.SetLogger(logs.AdapterFile,logPath)
	if err != nil{
		// do something
	}
	// 启动微信公众号后台服务
	wechat.Run()

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
