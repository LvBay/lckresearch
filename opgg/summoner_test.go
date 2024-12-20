package opgg

import (
	"log"
	"spider/service"
	"testing"
)

func TestSummonerRankCrawler_GetSummonerRank(t *testing.T) {
	s := &SummonerRankCrawler{}
	got := s.GetSummonerRank(service.GetSummonerRankReq{Start: 1, End: 3})
	log.Println(got)
}
