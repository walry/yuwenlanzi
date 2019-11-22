package menu

import "yuwenlanzi/common"

type NodeMenu struct {
	SubButton 				[]*NodeMenu			`json:"sub_button"`
	Name 					string 				`json:"name"`
	Type 					string 				`json:"type"`
	Key 					string 				`json:"key"`
}

func create(name string, t string, key string) *NodeMenu {
	return &NodeMenu{
		Name: name,
		Type: t,
		Key: key,
	}
}

// 创建自定义菜单

func GetDefineMenu() []*NodeMenu {
	menu1 := create("生活服务","click","live-service")
	menu1.SubButton = append(menu1.SubButton,
			create("新闻咨询","click",common.NEWS),
			create("笑话大全","click",common.JOKES),
			create("微信精选","click",common.WEXIN_SELECTION),
			//create("成语词典","click",common.DIRECTORY),
			//create("历史上的今天","click","live-service.history-today"),
			//create("万年历","click","live-service.perpetual-calendar"),
		)
	menu2 := create("查询","click","query")
	menu2.SubButton = append(menu2.SubButton,
			create("天气预报","click",common.WEATHER_REPORT),
			//create("QQ号码测吉凶","click",common.QQ_TEST),
			//create("汇率","click",common.EXCHANGE_RATE),
			//create("黄金数据","click",common.GOLD_DATA),
			create("彩票开奖","click","query.lottery"),
			create("星座运势","click","query.horoscope"),
		)
	menu3 := create("实用工具","click","tool")
	menu3.SubButton = append(menu3.SubButton,
			create("生成二维码","click",common.QR_CODE_CREATE),
			create("字体转换","click",common.WORD_EXCHANGE),
		)

	return []*NodeMenu{menu1,menu2,menu3}
}


