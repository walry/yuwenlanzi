package api

type News struct {
	Uniquekey 						string 							`json:"uniquekey"`
	Title 							string							`json:"title"`
	Date 							string 							`json:"date"`
	Category 						string							`json:"category"`
	AuthorName 						string							`json:"author_name"`
	Url 							string							`json:"url"`
	ThumbnailPicS 					string 							`json:"thumbnail_pic_s"`
	ThumbnailPicS02 				string							`json:"thumbnail_pic_s02"`
	ThumbnailPicS03 				string							`json:"thumbnail_pic_s03"`
}

type NewsResult struct {
	Stat 							string 							`json:"stat"`
	Data 							[]*News 						`json:"data"`
}


func PopNews() *News {
	if len(list) == 0 {
		jvHe := &JvHe{}
		jvHe.GetNews()
	}
	pop := list[0]
	list = list[1:]
	return pop
}

func AppendNews(data []*News) {
	for _, news := range data {
		list = append(list, news)
	}
}

