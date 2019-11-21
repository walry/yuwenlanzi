package api


var(
	//新闻咨询数组
	list 							[]*News
	//笑话数组
	jokes 							[]*Joke
	//微信精选
	selections 						[]*Selection
)

type JvHeResponse struct {
	ErrorCode 							int 					`json:"error_code"`
	Reason 								string 					`json:"reason"`
	Result 								interface{}				`json:"result"`
}
