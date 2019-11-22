package api

type JokeResult struct {
	Data 								[]*Joke
}

type Joke struct {
	Content 							string 					`json:"content"`
	HashId 								string 					`json:"hashId"`
	Unixtime 							int64 					`json:"unixtime"`
	UpdateTime 							string 					`json:"updatetime"`
}

func PopOneJoke() string {
	if len(jokes) == 0 {
		jh := &JvHe{}
		jh.GetJokes()
	}
	pop := jokes[0]
	jokes = jokes[1:]
	return pop.Content
}

func UpdateJokePool(data []*Joke){
	jokes = data
}

