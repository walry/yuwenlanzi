package wechat

import (
	"yuwenlanzi/service/wechat/menu"
	"yuwenlanzi/service/wechat/token"
)

func Run()  {

	// 获取access_token
	token.Start()
	//配置菜单
	menu.ConfigDIYMenu()
}