package lolps

import (
	"lckresearch/service"
	"testing"
)

func TestGetChampionItemPerformance(t *testing.T) {
	ins := LolPsChampionCrawler{}
	req := service.GetChampionItemPerformanceReq{
		ChampionId: 65,
	}
	ins.GetChampionItemPerformance(req)
}
