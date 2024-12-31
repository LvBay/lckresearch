package lolps

import (
	"context"
	"lckresearch/service"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
)

type LolPsComparer struct {
}

var CompareRes = make(map[string]int, 100)

func (l *LolPsComparer) Compare(matchInfo service.MatchInfo, perf service.ChampionItemPerformance) (diff bool) {
	// 比较对局出装和常规出装
	// 如果diff，则返回true
	items := fiterItems(matchInfo)
	if len(items) == 2 {
		diff = compareItems(items, perf.Two)
	} else if len(items) == 3 {
		diff = compareItems(items, perf.Three)
	} else if len(items) == 4 {
		diff = compareItems(items, perf.Four)
	}
	if diff {
		itemsStr := itemsToString(fiterItems(matchInfo))
		log.Println("发现新套路", getChampionName(matchInfo.ChampionId),
			"对局出装:", itemsStr)
		key := getChampionName(matchInfo.ChampionId) + "_" + itemsStr
		CompareRes[key]++
	}

	return false
}

func compareItems(items []service.Item, perfItems [][]service.ItemPerformance) (diff bool) {
	hasSame := false
	for _, tmpList := range perfItems {
		coreItems := make([]service.Item, 0, 2)
		for _, tmp := range tmpList {
			coreItems = append(coreItems, tmp.Item)
		}
		cKey := buildKey(coreItems)
		mKey := buildKey(items)
		if cKey == mKey {
			hasSame = true
			// log.Println("属于常规出装", "对局出装:", mKey, ToString(items))
			break
		}
	}
	if !hasSame {
		// log.Println("发现新套路 ", "对局出装:", itemsToString(items))
	}
	return !hasSame
}

func buildKey(items []service.Item) string {
	newItems := make([]service.Item, len(items))
	copy(newItems, items)
	sort.Slice(newItems, func(i, j int) bool {
		return newItems[i].ItemId < newItems[j].ItemId
	})
	str := make([]string, 0, len(newItems))
	for _, v := range newItems {
		str = append(str, strconv.Itoa(v.ItemId))
	}
	return strings.Join(str, "_")
}

func itemsToString(items []service.Item) string {
	newItems := make([]service.Item, len(items))
	copy(newItems, items)
	sort.Slice(newItems, func(i, j int) bool {
		return newItems[i].ItemId < newItems[j].ItemId
	})

	str := make([]string, 0, len(items))
	for _, v := range items {
		str = append(str, v.ItemName)
	}
	return strings.Join(str, "_")
}

func fiterItems(matchInfo service.MatchInfo) []service.Item {
	list := matchInfo.ItemList
	newList := make([]service.Item, 0, len(list))
	for _, item := range list {
		item := getItem(item.ItemId)
		if !item.IsCore || item.IsShoes || item.Price < 1900 {
			// log.Println("忽略价格低于1000的物品", item.Name)
			continue
		}
		newList = append(newList, service.Item{
			ItemId:   item.ItemId,
			Price:    item.Price,
			ItemName: item.NameCn,
		})
	}
	return newList
}

func GetBaseItems() {
	// fetchItemList()
}

func fetchItemList() {

	url := "https://game.gtimg.cn/images/lol/act/img/js/items/items.js"
	ret := dataClient.GetVar(context.Background(), url)
	mList := gjson.New(ret).Get("items").Maps()

	for _, v := range mList {
		item := ItemInfo{}
		gconv.Struct(v, &item)
		key := item.ItemId
		itemCache.Set(context.Background(), key, item, 0)
	}
}

type ItemInfo struct {
	// ItemId      string `json:"itemId"`
	// Name        string `json:"name"`
	// Total       string `json:"total"`
	// Description string `json:"description"`

	ItemId  int    `json:"itemId"`
	NameCn  string `json:"name"`
	Price   int    `json:"total"`
	IsCore  bool   `json:"isCore"`
	IsShoes bool   `json:"isShoes"`
	DescCn  string `json:"descCn"`
}

var itemCache = gcache.New()

var dataClient = gclient.New().Use(headerMid)

func headerMid(cli *gclient.Client, r *http.Request) (response *gclient.Response, err error) {
	r.Header.Set("user-agent", "curl/7.88.1")
	return cli.Next(r)
}
