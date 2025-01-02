package main

import (
	"lckresearch/lolps"
	"lckresearch/service"
	"log"
	"time"
)

func main() {
	lolps.InitRedis()
	lolps.GetBaseItems()
	lolps.InitChampionNames()
	ins := lolps.SummonerRankCrawler{}
	req := service.GetSummonerRankReq{}
	req.Region = "kr"
	req.Start = 16
	req.End = 17
	res := ins.GetSummonerRank(req)
	log.Println("ur:", res[0].URLPath)
	champIns := lolps.LolPsChampionCrawler{}
	cmIns := lolps.LolPsComparer{}
	now := time.Now()
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
	spend := time.Since(now)
	// 一共花费的秒数
	log.Println("spend:", spend, "平均", int(spend.Seconds())/(req.End-req.Start+1))
	log.Println("res:", lolps.ToString(lolps.CompareRes))
}
