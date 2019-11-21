package api


type LotteryList struct {
	LotteryId 							string 						`json:"lottery_id"`
	LotteryName 						string 						`json:"lottery_name"`
	LotteryTypeId 						string 						`json:"lottery_type_id"`
	Remarks 							string 						`json:"remarks"`
}

type LotteryInfo struct {
	LotteryId 							string 						`json:"lottery_id"`
	LotteryName 						string 						`json:"lottery_name"`
	LotteryRes 							string 						`json:"lottery_res"`
	LotteryNo 							string 						`json:"lottery_no"`
	LotteryDate 						string 						`json:"lottery_date"`
	LotteryExdate 						string 						`json:"lottery_exdate"`
	LotterySaleAmount 					string 						`json:"lottery_sale_amount"`
	LotteryPoolAmount 					string 						`json:"lottery_pool_amount"`
	LotteryPrize 						[]*LPrize 					`json:"lottery_prize"`
}

type LPrize struct {
	PrizeName 							string 						`json:"prize_name"`
	PrizeNum 							string 						`json:"prize_num"`
	PrizeAmount 						string 						`json:"prize_amount"`
	PrizeRequire 						string 						`json:"prize_require"`
}