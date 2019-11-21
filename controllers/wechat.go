package controllers

import (
	"crypto/sha1"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
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
	if err := xml.Unmarshal(we.Ctx.Input.RequestBody,&wm.Ctx); err !=nil {
		fmt.Print(err)
		logs.Info("xml.Unmarshal ----",err.Error())
		wm.WriteText("公众号出了点问题，开发人员会第一时间处理的！")
		return
	}
	logs.Info("wm.Ctx------",wm.Ctx)
	b,_ := json.Marshal(wm.Ctx)
	h := sha1.New()
	h.Write(b)
	wm.RequestId = fmt.Sprintf("%x",h.Sum(nil))
	logs.Info("wm.RequestId------",wm.RequestId)
	wm.Parse()
	logs.Info("wm.ResponseXml------",wm.ResponseXml)
	we.Data["xml"] = wm.ResponseXml[wm.RequestId]
	we.ServeXML()
}


//验证服务器
func (we *WechatController) Auth() {
	signature := we.GetString("signature")
	timestamp,_ := we.GetInt("timestamp")
	nonce,_ := we.GetInt("nonce")
	fmt.Print(signature,"\n",timestamp,"\n",nonce,"\n")
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
