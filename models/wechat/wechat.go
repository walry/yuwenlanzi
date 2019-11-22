package wechat

import (
	"encoding/xml"
	"time"
	"yuwenlanzi/common"
	"yuwenlanzi/tools/api"
)

type ChatModal struct {
	RequestId 						string
	RequestBody 					[]byte
	Ctx 							*BaseData
	ResponseXml 					map[string]interface{}
}

//解析推送到后台的消息,根据事件类型做不同处理
func (wm *ChatModal) Parse(){
	switch wm.Ctx.MsgType.Value {
		//处理事件消息
		case "event": wm.ParseEvent()
		//处理文本消息
		case "text": wm.HandleMessage()
	}
	//其他情况暂时回复空消息
	wm.WriteText("")
}

func (wm *ChatModal) ParseEvent(){
	var event EventData
	_ = xml.Unmarshal(wm.RequestBody,&event)

	if event.Event.Value == "CLICK" {
		switch event.EventKey.Value {

		case common.NEWS : wm.WriteNews()
		case common.JOKES : wm.WriteJokes()
		case common.LOTTERY_QUERY : wm.ResponseLottery()

		}
	}
}

func (wm *ChatModal) HandleMessage(){

	var text TextContent
	_ = xml.Unmarshal(wm.RequestBody,&text)
	switch text.Content.Value {

	case "查询彩票" : wm.ResponseLottery()
	case "看新闻"   : wm.WriteNews()
	case "笑话大全" : wm.WriteJokes()
	case "微信精选" : wm.ViewWechatSelection()


	default:
		wm.ChatWithRobot(text.Content.Value)
		
	}
}

func (wm *ChatModal) ResponseLottery(){
	jh := &api.JvHe{}
	info := jh.GetLotteryInfo()
	wm.WriteText(info)
}

//调用问答机器人和客户端交流
func (wm *ChatModal) ChatWithRobot(text string){
	jh := &api.JvHe{}
	message := jh.CallRobot(text,wm.RequestId)
	wm.WriteText(message)
}

//回复图文消息，点击菜单”新闻咨询“响应
func (wm *ChatModal) WriteNews() {
	news := api.PopNews()
	res := make(map[string]interface{})
	res[wm.RequestId] = &ImageTextResponse{
		BaseData: BaseData{
			FromUserName: wm.Ctx.ToUserName,
			ToUserName: wm.Ctx.FromUserName,
			CreateTime: time.Now().Unix(),
			MsgType: CDATA{ Value: "news" } ,
		},
		ArticleCount: 1,
		Articles: Articles{
			List: []Article{
				{
					Title: CDATA{
						Value:news.Title,
					},
					Description: CDATA{
						Value: "",
					},
					PicUrl: CDATA{
						Value: news.ThumbnailPicS,
					},
					Url:CDATA{
						Value: news.Url,
					},
				},
			},
		},
	}
	wm.ResponseXml = res
}

//回复文本消息
func (wm *ChatModal)WriteText(str string){
	res := make(map[string]interface{})
	res[wm.RequestId] = &TextResponse{
		BaseData:BaseData{
			ToUserName: wm.Ctx.FromUserName,
			FromUserName: wm.Ctx.ToUserName,
			CreateTime: time.Now().Unix(),
			MsgType:CDATA{
				Value: "text",
			},
		},
		Content: CDATA{
			Value: str,
		},
	}

	wm.ResponseXml = res
}

func (wm *ChatModal) WriteImageText(title string,description string,picSrc string,url string){
	res := make(map[string]interface{})
	res[wm.RequestId] = &ImageTextResponse{
		BaseData: BaseData{
			FromUserName: wm.Ctx.ToUserName,
			ToUserName: wm.Ctx.FromUserName,
			CreateTime: time.Now().Unix(),
			MsgType: CDATA{ Value: "news" } ,
		},
		ArticleCount: 1,
		Articles: Articles{
			List: []Article{
				{
					Title: CDATA{
						Value:title,
					},
					Description: CDATA{
						Value: description,
					},
					PicUrl: CDATA{
						Value: picSrc,
					},
					Url:CDATA{
						Value: url,
					},
				},
			},
		},
	}
	wm.ResponseXml = res
}

func (wm *ChatModal) WriteJokes() {
	joke := api.PopOneJoke()
	wm.WriteText(joke)
}

func (wm *ChatModal) ViewWechatSelection(){
	selection := api.PopOneSelection()
	if selection.FirstImg == "" {
		selection.FirstImg = wechatDefaultUrl
	}
	wm.WriteImageText(selection.Title,selection.Mark,selection.FirstImg,selection.Url)
}
