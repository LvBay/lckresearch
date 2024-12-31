package lolps

// https://lol.ps/ranking?lane=-1&page=1&region=kr

import (
	"encoding/json"
	"fmt"
	"lckresearch/service"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/go-rod/rod"
)

// 实现service.SummonerRankCrawler
type SummonerRankCrawler struct {
}

var localBrowser = rod.New().MustConnect()

// #content-container > div.divide-y.divide-psgr-formborder.border-y.border-psgr-formtext.dark\:divide-dark-psgr-formborder.dark\:border-dark-psgr-formtext.desktop\:h-full.desktop\:border-y-psbk-sub.dark\:desktop\:border-y-dark-psbk-sub > div:nth-child(1) > div.flex.h-full.w-\[calc\(100vw-144px\)\].grow.items-center.mobile-small\:w-\[calc\(100vw-184px\)\].desktop\:gap-2 > a.block.w-full.grow.truncate.text-sm.font-bold.hover\:underline.desktop\:w-\[168px\].desktop\:text-rg
// //*[@id="content-container"]/div[4]/div[1]/div[1]/a[1]
// /html/body/main/section/div[4]/div[2]/div[1]/i

func (s *SummonerRankCrawler) GetSummonerRank(req service.GetSummonerRankReq) []service.SummonerInfo {
	var summonerList = make([]service.SummonerInfo, 0, req.End-req.Start+1)

	region := req.Region
	pageStart := req.Start / 100
	if pageStart == 0 {
		pageStart = 1
	}
	pageEnd := req.End / 100
	if pageEnd >= 10 {
		pageEnd = 10
	}
	if pageStart > pageEnd {
		pageStart = pageEnd
	}
	count := req.End - req.Start + 1
	var err error
	var el *rod.Element
	for page := pageStart; page <= pageEnd; page++ {
		if len(summonerList) >= req.End-req.Start+1 {
			break
		}
		page := smartPage(fmt.Sprintf("https://lol.ps/ranking/?lane=-1&page=%d&region=%s", page, region))
		for i := 1; i <= 50; i++ {
			if count <= 0 {
				break
			}

			el, err = page.Timeout(2 * time.Second).ElementX(fmt.Sprintf(`//*[@id="content-container"]/div[4]/div[%d]/div[1]/a[1]`, i))
			if err != nil {
				str, _ := page.HTML()
				log.Printf("获取summoner 列表失败:%v", str)
				break
			}
			if el == nil {
				break
			}
			id := el.MustAttribute("href")
			u, _ := url.Parse(*id)
			sid := u.Query().Get("sid")
			summoner := service.SummonerInfo{
				URLPath: *id,
				ID:      sid,
				Name:    strings.Replace(strings.TrimPrefix(u.Path, "/summoner/"), "_", "#", -1),
			}
			summonerList = append(summonerList, summoner)
			count--
		}
	}
	log.Printf("抓取%s rank选手排行，数量:%d,id:%v\n", region, len(summonerList), ToString(summonerList))
	return summonerList
}

func ToString(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func smartPage(u string) *rod.Page {
	var page *rod.Page
	for i := 0; i < 3; i++ {
		page = localBrowser.MustPage(u)

		str := page.MustHTML()
		if strings.Contains(str[:200], "Just a moment") {
			log.Println(u, "检测到反爬虫，等待5s", i)
			localBrowser = rod.New().MustConnect()
			time.Sleep(5 * time.Second)
		} else {
			return page
		}
	}
	return page
}
