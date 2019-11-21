package api

import "math"

type Selection struct {
	Id 							string 							`json:"id"`
	Title 						string 							`json:"title"`
	Source 						string 							`json:"source"`
	FirstImg 					string 							`json:"FirstImg"`
	Mark 						string 							`json:"mark"`
	Url 						string 							`json:"url"`
}

type SelectionResult struct {
	List 						[]*Selection 					`json:"list"`
	TotalPage 					int 							`json:"totalPage"`
	Ps 							int 							`json:"ps"`
	Pno 						int 							`json:"pno"`
}

var (
	selectionResultIndex int
)

func PopOneSelection() *Selection {
	if len(selections) == 0 {
		jh := &JvHe{}
		jh.GetWechatSelection(50,selectionResultIndex)
	}
	pop := selections[0]
	selections = selections[1:]
	return pop
}

func ReceiveNewSelection(result *SelectionResult){
	selections = result.List
	if max := math.Ceil(float64(result.Pno/50)) ; int(max) < selectionResultIndex + 1 {
		selectionResultIndex = 1
	}else {
		selectionResultIndex += 1
	}
}