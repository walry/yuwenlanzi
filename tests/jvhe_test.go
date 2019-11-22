package test

import (
	"fmt"
	"testing"
	"time"
	"yuwenlanzi/tools/api"
)

func TestGetNews(t *testing.T) {

	for i:= 0; i < 5; i++  {
		news := api.PopNews()
		fmt.Printf("news ---%+v\n",news)
		time.Sleep(2 * time.Second)
	}
}

func TestWriteNews(t *testing.T){
	//b := wechat.BaseData{
	//	ToUserName:wechat.CDATA{
	//		Value: "duzhengwei",
	//	},
	//	FromUserName:wechat.CDATA{
	//		Value: "yuwenlanzi",
	//	},
	//	CreateTime: time.Now().Unix(),
	//	MsgType:wechat.CDATA{
	//		Value: "event",
	//	},
	//}
	//eve := wechat.Event{
	//	BaseData: b,
	//	Event:wechat.CDATA{
	//		Value:"CLICK",
	//	},
	//	EventKey:wechat.CDATA{
	//		Value: "live-service.news",
	//	},
	//}
	//t.Log("--------------",eve)
	//em := wechat.EventModal{
	//	Eve: &eve,
	//}
	//
	//em.ParseEvent()
	//
	//t.Log(em.ResponseContent)
}

func TestLottery(t *testing.T){
	jh := &api.JvHe{}
	result := jh.GetLotteryInfo()
	t.Log(result)
}