package lolps

import (
	"lckresearch/service"
	"testing"
)

func TestGetChampionItemPerformance(t *testing.T) {
	InitRedis()
	ins := LolPsChampionCrawler{}
	req := service.GetChampionItemPerformanceReq{
		ChampionId: 112,
	}
	ins.GetChampionItemPerformance(req)
}
