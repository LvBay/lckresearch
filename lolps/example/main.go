package main

import (
	"lckresearch/lolps"
	"lckresearch/service"
	"log"
)

func main() {
	lolps.GetBaseItems()
	lolps.InitChampionNames()
	ins := lolps.SummonerRankCrawler{}
	req := service.GetSummonerRankReq{}
	req.Region = "kr"
	req.Start = 1
	req.End = 5
	res := ins.GetSummonerRank(req)
	log.Println("ur:", res[0].URLPath)
	champIns := lolps.LolPsChampionCrawler{}
	cmIns := lolps.LolPsComparer{}

	{
		for i, summoner := range res {
			log.Println("第", i+1, "个召唤师")
			ins := lolps.LolPsMatchCrawler{}
			req := service.GetMatchListReq{}
			req.SummonerInfo.URLPath = summoner.URLPath
			matchList, _ := ins.GetMatchList(req)
			{
				for j, match := range matchList {
					log.Println("第", j+1, "个比赛")
					champReq := service.GetChampionItemPerformanceReq{
						ChampionId: match.ChampionId,
					}
					perf, _ := champIns.GetChampionItemPerformance(champReq)
					cmIns.Compare(match, perf)
				}
			}
		}
	}
	log.Println("res:", lolps.ToString(lolps.CompareRes))
}
