package token

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"sync"
	"time"
)

var (
	requestUrl = "https://api.weixin.qq.com/cgi-bin/token"
	appId = beego.AppConfig.String("appid")
	secret = beego.AppConfig.String("secret")
)

const (
	EXPIRES = 7200
)

type WechatAccessToken struct {
	AccessToken 			string
	expiresIn 				float64
}
//定时刷新
func (wa *WechatAccessToken) Run(){
	ticker := time.NewTicker(EXPIRES * time.Second)
	for {
		select {
		case <-ticker.C:
			expires := wa.requestAccessToken()
			if expires > 0 {
				//重置过期时间
				ticker = time.NewTicker(time.Duration(expires) * time.Second)
			}
			break
		}
	}
}
//调用获取access_token接口
func (wa *WechatAccessToken) requestAccessToken() float64 {
	req := httplib.Get(requestUrl)
	req.Param("grant_type","client_credential")
	req.Param("appid",appId)
	req.Param("secret",secret)
	result := make(map[string]interface{})
	err := req.ToJSON(&result)
	if err != nil {
		fmt.Printf("Error happen %+v",err)
		return 0
	}

	if result["access_token"] != nil {
		m := sync.Mutex{}
		m.Lock()
		wa.AccessToken = result["access_token"].(string)
		wa.expiresIn = result["expires_in"].(float64)
		m.Unlock()
		return wa.expiresIn
	}else {
		fmt.Printf("response return error:%+v",result)
		return 0
	}
}
//提供被动刷新接口
func (wa *WechatAccessToken) Refresh(){
	wa.requestAccessToken()
}

//获取access_token
func GetAccessToken() string {
	wa := &WechatAccessToken{}
	if wa.AccessToken == "" {
		wa.Refresh()
	}
	return wa.AccessToken
}

func Start(){
	wa := WechatAccessToken{}
	go wa.Run()
}


