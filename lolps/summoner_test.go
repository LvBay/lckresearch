package lolps

import (
	"lckresearch/service"
	"testing"
)

func TestGetSummonerRank(t *testing.T) {
	ins := SummonerRankCrawler{}
	req := service.GetSummonerRankReq{}
	req.Region = "kr"
	req.Start = 1
	req.End = 3
	ins.GetSummonerRank(req)
}
