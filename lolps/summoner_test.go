package lolps

import (
	"lckresearch/service"
	"log"
	"testing"
)

func TestGetSummonerRank(t *testing.T) {
	ins := SummonerRankCrawler{}
	req := service.GetSummonerRankReq{}
	req.Region = "kr"
	req.Start = 1
	req.End = 3
	res := ins.GetSummonerRank(req)
	log.Println("ur:", res[0].URLPath)
	champIns := LolPsChampionCrawler{}
	cmIns := LolPsComparer{}

	{
		for _, summoner := range res {
			ins := LolPsMatchCrawler{}
			req := service.GetMatchListReq{}
			req.SummonerInfo.URLPath = summoner.URLPath
			matchList, _ := ins.GetMatchList(req)
			{
				for _, match := range matchList {
					champReq := service.GetChampionItemPerformanceReq{
						ChampionId: match.ChampionId,
					}
					perf, _ := champIns.GetChampionItemPerformance(champReq)
					cmIns.Compare(match, perf)
				}
			}
		}
	}
}
