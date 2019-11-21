package wechat

import "encoding/xml"




var (
	wechatDefaultUrl = "https://mmbiz.qpic.cn/mmbiz_png/K46olEVMhnu59HrZr33pF4qBVjKMQCaDT2V1kh898dpCUKb23QPv8VTJUQeicdfVfI6qqmJZrXopfyBBRRUWLkw/0?wx_fmt=png"
)


type CDATA struct {
	Value 				string 				`xml:",cdata"`
}

type GlobalData struct {
	BaseData

	//推送事件
	Event 						CDATA 					`xml:"Event"`
	EventKey 					CDATA 					`xml:"EventKey"`
	Ticket 						CDATA 					`xml:"Ticket"`

	//上报地理位置事件
	Latitude 					CDATA 					`xml:"Latitude"`
	Longitude 					CDATA 					`xml:"Longitude"`
	Precision 					CDATA 					`xml:"Precision"`

	Content 					CDATA 					`xml:"Content"`
	MsgId 						int64 					`xml:"MsgId"`

}

type BaseData struct {
	ToUserName 					CDATA 					`xml:"ToUserName"`
	FromUserName 				CDATA 					`xml:"FromUserName"`
	CreateTime 					int64 					`xml:"CreateTime"`
	MsgType 					CDATA 					`xml:"MsgType"`
}

//图文消息模板
type ImageTextResponse struct {
	XMLName 					xml.Name 					`xml:"xml"`
	BaseData
	ArticleCount 				int64 						`xml:"ArticleCount"`
	Articles 					Articles					`xml:"Articles"`
}

type Articles struct {
	XMLName 					xml.Name 					`xml:"Articles"`
	List 						[]Article 					`xml:"item"`
}

type Article struct {
	XMLName 					xml.Name 					`xml:"item"`
	Title 						CDATA 						`xml:"Title"`
	Description 				CDATA 						`xml:"Description"`
	PicUrl 						CDATA 						`xml:"PicUrl"`
	Url 						CDATA 						`xml:"Url"`
}

type TextResponse struct {
	XMLName 					xml.Name 				`xml:"xml"`
	BaseData
	Content 					CDATA 					`xml:"Content"`
	MsgId 						int64 					`xml:"MsgId"`
}