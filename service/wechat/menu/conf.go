package menu

import "yuwenlanzi/service/wechat/token"

const (
	QUERY_MENU_INFO = "https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info"
	CREATE_MENU = "https://api.weixin.qq.com/cgi-bin/menu/create"
)

var (
	accessToken = token.GetAccessToken()
)

