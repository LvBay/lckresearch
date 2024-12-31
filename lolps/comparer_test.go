package lolps

import (
	"fmt"
	"lckresearch/service"
	"log"
	"testing"
)

func TestCompare(t *testing.T) {
	log.Println("11")
	GetBaseItems()
	InitChampionNames()
	ins := LolPsMatchCrawler{}
	req := service.GetMatchListReq{}
	req.SummonerInfo.URLPath = "/summoner/민철이여친구함_0415?sid=4txgTW5VnMA2NhHKwrQu1L-2exTQvUjmL_-ADqSBILH9po0&region=kr"
	matchList, _ := ins.GetMatchList(req)

	champIns := LolPsChampionCrawler{}

	cm := LolPsComparer{}
	for _, match := range matchList {
		champReq := service.GetChampionItemPerformanceReq{
			ChampionId: match.ChampionId,
		}
		perf, _ := champIns.GetChampionItemPerformance(champReq)
		cm.Compare(match, perf)
	}
}

func TestGetItem(t *testing.T) {
	item := getItem(3157)
	fmt.Println(item)
}
