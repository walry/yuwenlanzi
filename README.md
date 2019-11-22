# 公众号后台项目
为学习Golang而起的练习项目，基于[Beego](https://beego.me/)框架,调用[聚合数据](https://www.juhe.cn/) API实现了当期**彩票查询**
、**查看新闻咨询**、**笑话大全**、**微信精选**以及**问答机器人**的功能，都是些比较简单的小功能，这个项目的目的就是熟悉下**微信公众号开发**流程、Beego框架使用和Golang语法学习，所以越简单越容易学习，有感兴趣的同学欢迎交流！

* 配置文件
```
appname = yuwenlanzi
httpport = 8080
runmode = prod
autorender = false
copyrequestbody = true
EnableDocs = true

appid=公众号后台能找到
secret=也在公众号后台里

logpath=/var/log/yuwenlanzi
```

* 验证微信服务器
```
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
```
* 部署脚本
```$xslt
#! /bin/bash

if [ $1 == "prod" ]; then
    ps -ef | grep -w yuwenlanzi | awk '{print $2}' | xargs kill -9
    git checkout master
    git fetch origin master
    git merge FETCH_HEAD
    cp ./conf/prod.app.conf ./conf/app.conf
    go install
    go build -o yuwenlanzi main.go
    nohup ./yuwenlanzi > /dev/null 2>&1 &
fi

if [ $1 == "dev" ]; then
    ps -ef | grep -w yuwenlanzi-dev | awk '{print $2}' | xargs kill -9
    git checkout dev
    git fetch origin dev
    git merge FETCH_HEAD
    cp ./conf/dev.app.conf ./conf/app.conf
    go install
    go build -o yuwenlanzi-dev main.go
    nohup ./yuwenlanzi-dev > /dev/null 2>&1 &
fi

exit
```
