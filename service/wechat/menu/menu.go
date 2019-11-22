package menu

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
)

type Menu struct {
	IsMenuOpen 					float64 // 自定义菜单是否开启
	IsUpdateMenu 				bool  //是否需要更新菜单
}

//获取配置信息
func (m *Menu) GetMenuInfo() {
	r := httplib.Get(QUERY_MENU_INFO)
	r.Param("access_token",accessToken)
	data := make(map[string]interface{})
	err := r.ToJSON(&data)
	if err != nil {
		logs.Info("getMenuInfo error:%+v",err)
		return
	}
	m.IsMenuOpen = data["is_menu_open"].(float64)
}

//自定义菜单
func (m *Menu) createMenu() {
	req := httplib.Post(CREATE_MENU + "?access_token=" + accessToken)
	menuArr := GetDefineMenu()
	b := make(map[string]interface{})
	b["button"] = menuArr
	r, _ := req.JSONBody(b)
	res := make(map[string]interface{})
	e := r.ToJSON(&res)
	if e != nil {
		fmt.Printf("define menu error:%+v \n",e)
		return
	}
	logs.Info("createMenu result-----",res)
}
//更新菜单
func (m *Menu) Update() {
	if m.IsUpdateMenu {
		m.createMenu()
		m.IsUpdateMenu = false
	}
}

func ConfigDIYMenu(){
	menu := &Menu{}
	menu.IsUpdateMenu = false
	menu.Update()
}

