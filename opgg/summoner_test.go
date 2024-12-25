package opgg

import (
	"log"
	"spider/service"
	"testing"
)

func TestSummonerRankCrawler_GetSummonerRank(t *testing.T) {
	s := &SummonerRankCrawler{}
	req := service.GetSummonerRankReq{Start: 1, End: 3}
	req.Region = "kr"
	got := s.GetSummonerRank(req)
	log.Println(got)
}
