package service

type MatchInfo struct {
	MatchId      string
	ChampionId   int
	SummonerName string
	ItemList     []Item
	GameTime     int
}

type GetMatchListReq struct {
	CommonParam
	SummonerInfo
	StartTime int64
	EndTime   int64
}

type MatchCrawler interface {
	GetMatchList(req GetMatchListReq) ([]MatchInfo, error)
	GetMatchInfo(matchId string) (MatchInfo, error)
}
