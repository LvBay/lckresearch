package opgg

import (
	"fmt"
	"log"
	"net/url"
	"spider/service"
	"strings"

	"github.com/gocolly/colly"
)

// 实现service.SummonerRankCrawler
type SummonerRankCrawler struct {
}

func (s *SummonerRankCrawler) GetSummonerRank(req service.GetSummonerRankReq) []string {
	var summonerList = make([]string, 0, req.End-req.Start+1)

	rankCollector := colly.NewCollector()
	stop := false
	rankCollector.OnHTML("tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(i int, trEl *colly.HTMLElement) {
			trEl.ForEach("td .summoner-link", func(j int, tdEl *colly.HTMLElement) {
				href := tdEl.Attr("href")
				hList := strings.Split(href, "/")
				name := hList[len(hList)-1]
				name, _ = url.PathUnescape(name)
				// log.Println("name:", i, name)
				if len(summonerList) >= req.End-req.Start+1 {
					stop = true
					return
				}
				summonerList = append(summonerList, name)
			})
		})
	})
	region := req.Region
	pageStart := req.Start / 100
	if pageStart == 0 {
		pageStart = 1
	}
	pageEnd := req.End / 100
	if pageEnd == 0 {
		pageEnd = 10
	}
	if pageEnd >= 10 {
		pageEnd = 10
	}
	if pageStart > pageEnd {
		pageStart = pageEnd
	}
	for i := pageStart; i <= pageEnd; i++ {
		if stop {
			break
		}
		u := fmt.Sprintf("https://www.op.gg/leaderboards/tier?region=%s&type=ladder&page=%d", region, i)
		log.Println("访问:", u)
		rankCollector.Visit(u)
	}
	log.Printf("抓取%s rank选手排行，数量:%d\n", region, len(summonerList))
	return summonerList
}
