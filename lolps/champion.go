package lolps

import (
	"context"
	"fmt"
	"lckresearch/service"
	"log"
	"time"

	"github.com/go-rod/rod"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
)

type LolPsChampionCrawler struct {
}

var champCache = gcache.New()

func (l *LolPsChampionCrawler) GetChampionItemPerformance(req service.GetChampionItemPerformanceReq) (service.ChampionItemPerformance, error) {
	val, _ := champCache.GetOrSetFunc(context.Background(), gconv.String(req), func(ctx context.Context) (value interface{}, err error) {
		u := fmt.Sprintf("https://lol.ps/champ/%d", req.ChampionId)
		page := smartPage(u)
		time.Sleep(3 * time.Second)
		log.Println("page load", u)
		var el *rod.Element
		for i := 0; i < 3; i++ {
			el, err = page.Timeout(3 * time.Second).ElementX(`//*[@id="content-container"]/div[1]/section[7]/div/div[1]`)
			if err == nil {
				break
			}
			log.Println("retry...", i)
		}
		if err != nil {
			str, _ := page.HTML()
			log.Println("finish retry", err, str)
			return service.ChampionItemPerformance{}, nil
		}
		els := el.MustElements("li")
		log.Println("els:", len(els))

		resp := service.ChampionItemPerformance{}

		resp.ChampionId = req.ChampionId
		resp.Two = coreItem(`//*[@id="content-container"]/div[1]/section[7]/div/div[1]`, page)
		resp.Three = coreItem(`//*[@id="content-container"]/div[1]/section[7]/div/div[2]`, page)
		resp.Four = coreItem(`//*[@id="content-container"]/div[1]/section[7]/div/div[3]`, page)
		return resp, nil
	}, 0)
	return val.Interface().(service.ChampionItemPerformance), nil

}

func coreItem(x string, page *rod.Page) [][]service.ItemPerformance {
	el := page.MustElementX(x)
	els := el.MustElements("li")

	itemsList := make([][]service.ItemPerformance, 0, 10)
	for _, el := range els {
		imgEls := el.MustElements("img")
		items := make([]service.ItemPerformance, 0, 4)
		for _, img := range imgEls {
			id := getItemIdFromImgSrc(img)
			tmp := service.ItemPerformance{}
			tmp.Item = service.Item{ItemId: gconv.Int(id)}
			items = append(items, tmp)
		}
		// log.Println("items:", items)
		itemsList = append(itemsList, items)
	}
	return itemsList
}

// 完成了英雄核心装备的爬取
// TODO 比对实际出装和常规出装
