package lolps

import (
	"lckresearch/service"
	"testing"
)

func TestGetMatchList(t *testing.T) {
	ins := LolPsMatchCrawler{}
	req := service.GetMatchListReq{}
	req.SummonerInfo.URLPath = "/summoner/민철이여친구함_0415?sid=4txgTW5VnMA2NhHKwrQu1L-2exTQvUjmL_-ADqSBILH9po0&region=kr"
	ins.GetMatchList(req)
}
