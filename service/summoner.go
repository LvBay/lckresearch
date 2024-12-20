package service

type CommonParam struct {
	Region      string
	Queue       string
	Tier        string
	GameVersion string
}

type GetSummonerRankReq struct {
	CommonParam
	Start int
	End   int
}

type SummonerRankCrawler interface {
	// 获取召唤师排名
	GetSummonerRank(req GetSummonerRankReq) []string
}
