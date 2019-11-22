package api

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"math/rand"
	"strconv"
	"time"
	"yuwenlanzi/common"
)

type JvHe struct {}

const (
	NEWS_URL = "http://v.juhe.cn/toutiao/index"
	JOKES_URL = "http://v.juhe.cn/joke/content/list.php"
	SELECTION_URL = "http://v.juhe.cn/weixin/query"
	ROBOT_URL = "http://op.juhe.cn/robot/index"
	LOTTERY_LIST_URL = "http://apis.juhe.cn/lottery/types"
	LOTTERY_QUERY_URL = "http://apis.juhe.cn/lottery/query"
 )

var (
	newsType = [10]string{"top","shehui","guonei","guoji","yule","tiyu","junshi","keji","shishang","caijing"}
)

//获取新闻
func (jh *JvHe) GetNews(){
	q := httplib.Get(NEWS_URL)
	q.Param("key",common.NEWS_KEY)
	rand.Seed(time.Now().UnixNano())
	q.Param("type",newsType[rand.Intn(10)])

	var result json.RawMessage
	res := JvHeResponse{
		Result: &result,
	}
	err := q.ToJSON(&res)

	if err != nil {
		fmt.Printf("request news error:%+v\n",err)
		return
	}

	if res.Reason == "成功的返回" {
		var newsResult NewsResult
		if e := json.Unmarshal(result,&newsResult); e != nil{
			logs.Error(e)
		}
		AppendNews(newsResult.Data)
	}
}

//获取笑话
func (jh *JvHe) GetJokes(){
	q := httplib.Get(JOKES_URL)
	q.Param("key",common.JOKES_KEY)
	q.Param("time",strconv.FormatInt(time.Now().Unix(),10))
	q.Param("sort","desc")
	q.Param("pagesize","20")
	q.Param("page","20")

	var joke json.RawMessage
	result := JvHeResponse{
		Result: &joke ,
	}
	err := q.ToJSON(&result)
	if err != nil {
		logs.Info("request jokes error:%+v\n",err)
		fmt.Printf("request jokes error:%+v\n",err)
		return
	}
	if result.ErrorCode == 0 {
		var j JokeResult
		_ = json.Unmarshal(joke,&j)
		UpdateJokePool(j.Data)
	}
}

//获取微信精选
func (jh *JvHe) GetWechatSelection(size int,index int){
	q := httplib.Get(SELECTION_URL)
	q.Param("key",common.SELECTION_KEY)
	q.Param("ps",strconv.Itoa(size))
	q.Param("pno",strconv.Itoa(index))

	var selection json.RawMessage
	result := JvHeResponse{
		Result: &selection,
	}
	err := q.ToJSON(&result)
	if err != nil {
		logs.Info("request selections error:%+v\n",err)
		fmt.Printf("request selections error:%+v\n",err)
		return
	}
	if result.ErrorCode == 0 {
		var s SelectionResult
		_ = json.Unmarshal(selection,&s)
		ReceiveNewSelection(&s)
	}
}

//问答机器人
func (jh *JvHe) CallRobot(msg string,uid string) string{
	q := httplib.Get(ROBOT_URL)
	q.Param("key",common.ROBOT_KEY)
	q.Param("info",msg)
	q.Param("userid",uid)

	var data struct{
		Code 				int 				`json:"code"`
		Text 				string 				`json:"text"`
	}

	var result json.RawMessage
	response := &JvHeResponse{
		Result: &result,
	}

	if err := q.ToJSON(&response); err != nil {
		logs.Error(err)
		return err.Error()
	}
	if response.ErrorCode != 0 {
		return response.Reason
	}
	_ = json.Unmarshal(result,&data)
	return data.Text
}

func (jh *JvHe) GetLotteryInfo() string{
	q := httplib.Get(LOTTERY_LIST_URL)
	q.Param("key",common.LOTTERY_KEY)

	var lotteryList json.RawMessage
	res := JvHeResponse{
		Result: &lotteryList,
	}
	if err := q.ToJSON(&res); err != nil {
		logs.Error(err)
		return err.Error()
	}
	if res.ErrorCode == 0 {
		var list []*LotteryList
		_ = json.Unmarshal(lotteryList,&list)
		l := len(list)
		ch := make(chan map[string]interface{},l)
		count := make(chan int)
		lotteryProducer(list,ch,count)
		var detail string
		for i := 0; i < l; i++ {
			select {
			case <-count:
				info := <-ch
				detail += jh.HandleLotteryInfo(info) + "\n\n"
			}
		}
		return detail
	}
	return res.Reason
}

func (jh *JvHe) HandleLotteryInfo(info map[string]interface{}) string {

	m := map[string]string{"1":"福利彩票","2":"体育彩票"}
	lottery := info["lottery"].(*LotteryInfo)
	group := info["group"].(*LotteryList)
	format := "第%s期【%s】【%s】彩票开奖结果：\n\n\t%s\n\n开奖日期：%s\n截止兑换日期：%s\n"
	str := fmt.Sprintf(format,lottery.LotteryNo,m[group.LotteryTypeId],lottery.LotteryName,lottery.LotteryRes,lottery.LotteryDate,lottery.LotteryExdate)
	if lottery.LotterySaleAmount != "" {
		str = fmt.Sprintf("%s本期销售额：%s元\n",str,lottery.LotterySaleAmount)
	}
	if lottery.LotteryPoolAmount != "" {
		str = fmt.Sprintf("%s奖池滚存：%s元\n",str,lottery.LotteryPoolAmount)
	}
	if lottery.LotteryPrize != nil {
		var prizeStr string
		for _,prize := range lottery.LotteryPrize {

			prizeStr += fmt.Sprintf("【%s】:\n中奖数量：%s\n中奖金额：%s元\n中奖条件：%s\n",prize.PrizeName,prize.PrizeNum,prize.PrizeAmount,prize.PrizeRequire)
		}
		str = fmt.Sprintf("%s\n奖项：\n%s",str,prizeStr)
	}
	return str
}

func lotteryProducer(list []*LotteryList,ch chan<- map[string]interface{},n chan<- int) {
	for _,item := range list{
		go func(item *LotteryList) {
			q := httplib.Get(LOTTERY_QUERY_URL)
			q.Param("key",common.LOTTERY_KEY)
			q.Param("lottery_id",item.LotteryId)

			var lotteryInfo json.RawMessage
			res := JvHeResponse{
				Result: &lotteryInfo,
			}

			if err := q.ToJSON(&res); err != nil {
				logs.Error(err)
				n <- 1
				return
			}

			if res.ErrorCode == 0 {
				var lottery LotteryInfo
				_ = json.Unmarshal(lotteryInfo,&lottery)
				out := make(map[string]interface{})
				out["lottery"] = &lottery
				out["group"] = item
				ch<- out
			}
			n <- 1
		}(item)
	}
}

func (jh *JvHe) Dictionary() {

}


