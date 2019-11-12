package tools

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"time"
)

var (
	requestUrl = "https://api.weixin.qq.com/cgi-bin/token"
)

const (
	APPID = "wx9454af5745ea0028"
	SECRET = "51c83c6abd7c2872e3f467489d3b6a80"
	EXPIRES = 4
)

type WechatAccessToken struct {
	AccessToken 			string
	ExpiresIn 				float64
}
//定时刷新
func (wa *WechatAccessToken) Run(){
	ticker := time.NewTicker(EXPIRES * time.Second)
	for {
		select {
		case <-ticker.C:
			expires := wa.requestAccessToken()
			fmt.Print("time out\n",time.Now(),"\n")
			if expires > 0 {
				//重置过期时间
				ticker = time.NewTicker(time.Duration(expires) * time.Second)
				fmt.Print("--expires--",expires,"\n")
			}
			break
		}
	}
}
//调用获取access_token接口
func (wa *WechatAccessToken) requestAccessToken() float64 {
	req := httplib.Get(requestUrl)
	req.Param("grant_type","client_credential")
	req.Param("appid",APPID)
	req.Param("secret",SECRET)
	result := make(map[string]interface{})
	err := req.ToJSON(&result)
	if err != nil {
		fmt.Printf("Error happen %+v",err)
		return 0
	}

	if result["access_token"] != nil {
		wa.AccessToken = result["access_token"].(string)
		wa.ExpiresIn = result["expires_in"].(float64)
		return wa.ExpiresIn
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
func (wa *WechatAccessToken) getAccessToken() string {
	if wa.AccessToken == "" {
		wa.requestAccessToken()
	}
	return wa.AccessToken
}

func Start(){
	fmt.Print("wechat_access_token run")
	wa := WechatAccessToken{}
	go wa.Run()
}


