package controllers

import (
	"crypto/sha1"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"sort"
	"strconv"
	"yuwenlanzi/models/wechat"
)

type WechatController struct {
	beego.Controller
}

//接收来自腾讯服务器的请求
func (we *WechatController) Index(){
	if false {
		we.Auth()
		return
	}

	wm := &wechat.ChatModal{}
	wm.RequestBody = we.Ctx.Input.RequestBody
	var base wechat.BaseData
	_ = xml.Unmarshal(we.Ctx.Input.RequestBody,&base)
	wm.Ctx = &base
	b,_ := json.Marshal(wm.RequestBody)
	h := sha1.New()
	h.Write(b)
	wm.RequestId = fmt.Sprintf("%x",h.Sum(nil))
	wm.Parse()
	we.Data["xml"] = wm.ResponseXml[wm.RequestId]
	we.ServeXML()
}


//验证服务器
func (we *WechatController)Auth() {
	//接收微信服务器发来的参数
	signature := we.GetString("signature")
	timestamp,_ := we.GetInt("timestamp")
	nonce,_ := we.GetInt("nonce")
	token := "yuwenlanzi"

	arr := []string{ token, strconv.Itoa(timestamp), strconv.Itoa(nonce) }
	sort.Strings(arr)

	var tmpStr string
	for _,item := range arr{
		tmpStr += item
	}
	h := sha1.New()
	h.Write([]byte(tmpStr))
	str := fmt.Sprintf("%x",h.Sum(nil))
	h.Sum(nil)
	if str == signature {
		we.Ctx.WriteString(we.GetString("echostr"))
	}
}
