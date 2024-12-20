package service

type MatchInfo struct {
	MatchId      string
	ChampionId   int
	SummonerName string
	ItemList     []Item
}

type GetMatchListReq struct {
	CommonParam
	SummonerId string
	StartTime  int64
	EndTime    int64
}

type MatchCrawler interface {
	GetMatchList(req GetMatchListReq) ([]MatchInfo, error)
	GetMatchInfo(matchId string) (MatchInfo, error)
}
