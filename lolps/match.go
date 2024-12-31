package lolps

import (
	"context"
	"fmt"
	"lckresearch/service"
	"log"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/gogf/gf/v2/util/gconv"
)

type LolPsMatchCrawler struct {
}

// 获取比赛信息 实现service.MatchCrawler
func (l *LolPsMatchCrawler) GetMatchList(req service.GetMatchListReq) ([]service.MatchInfo, error) {
	u := "https://lol.ps" + req.URLPath
	log.Println("u:", u)
	page := smartPage(u)

	time.Sleep(2 * time.Second)
	log.Println("sleep 5s")

	// 模拟点击下一页，加载更多比赛数据
	for i := 0; i < 3; i++ {
		x := `//*[@id="content-container"]/div[2]/section/div/div[2]/div[3]/button`
		el := page.MustElementX(x)
		h := el.MustHTML()
		log.Println("html:", h)
		el.MustClick()
		time.Sleep(2 * time.Second)
	}

	matchList := make([]service.MatchInfo, 0, 60)

	for line := 1; line <= 50; line++ {
		items := make([]service.Item, 0, 7)
		match := service.MatchInfo{}
		// 英雄
		championImgEl := page.MustElementX(fmt.Sprintf(`//*[@id="content-container"]/div[2]/section/div/div[2]/div[2]/div[%d]/div/div[1]/div/div[2]/div[1]/div[1]/a/img`, line))
		match.ChampionId = gconv.Int(getItemIdFromImgSrc(championImgEl))

		// 装备
		for itemIdx := 1; itemIdx <= 7; itemIdx++ {
			x := fmt.Sprintf(`//*[@id="content-container"]/div[2]/section/div/div[2]/div[2]/div[%d]/div/div[1]/div/div[2]/div[1]/div[2]/div[%d]/div/div/img`, line, itemIdx)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			pageWithCancel := page.Context(ctx)
			go func() {
				time.Sleep(500 * time.Millisecond)
				cancel()
			}()
			el, err := pageWithCancel.ElementX(x)
			if err != nil {
				// log.Println("err", err)
				continue
			}
			id := getItemIdFromImgSrc(el)
			items = append(items, service.Item{ItemId: gconv.Int(id)})
		}
		match.ItemList = items
		match.SummonerName = req.SummonerInfo.Name

		matchList = append(matchList, match)
	}
	log.Println("matchList:", len(matchList))

	return matchList, nil
}

func getItemIdFromImgSrc(el *rod.Element) string {
	src := getImgElSrc(el)
	id := findItemIdFromImgSrc(src)
	return id
}

func getImgElSrc(el *rod.Element) string {
	src := el.MustAttribute("src")
	return *src
}

func findItemIdFromImgSrc(src string) string {
	start := strings.LastIndex(src, "/") + 1
	end := strings.LastIndex(src, "_")
	if end == -1 {
		return src
	}
	if start >= end {
		return src
	}
	return src[start:end]
}
