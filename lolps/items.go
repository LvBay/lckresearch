package lolps

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/util/gconv"
)

var champNameCache = make(map[int]string, 200)

type ChampName struct {
	NameCn     string
	ChampionId int
}

func InitChampionNames() {
	url := "https://lol.ps/api/info/champion-names"
	val := dataClient.GetVar(context.Background(), url)
	list := []ChampName{}
	gjson.New(val).Get("data").Struct(&list)
	for _, v := range list {
		champNameCache[v.ChampionId] = v.NameCn
	}
}

func getChampionName(id int) string {
	return champNameCache[id]
}

func getItem(id int) ItemInfo {
	itemVal, _ := itemCache.GetOrSetFunc(context.Background(), "item::"+strconv.Itoa(id), func(ctx context.Context) (value interface{}, err error) {
		url := fmt.Sprintf("https://lol.ps/api/info/item-info/%d", id)
		val := dataClient.GetVar(context.Background(), url)
		item := ItemInfo{}
		gjson.New(val).Get("data").Struct(&item)
		return item, nil
	}, time.Hour*6)
	ret := ItemInfo{}
	gconv.Struct(itemVal, &ret)
	return ret
}

// TODO
// 支持英雄+位置的筛选，例如上单亚索
// 装备进化：魔宗/魔切，凛冬之临/末日寒冬
// 支持筛选：时间、大区、对线
