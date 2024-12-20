package service

import "fmt"

type Comparer struct{}

type Item struct {
	ItemId   int
	ItemName string
	Price    int
}

func (c *Comparer) Compare(matchInfo MatchInfo, perf ChampionItemPerformance) (diff bool) {
	// 比较对局出装和常规出装
	// 如果diff，则返回true
	return false
}

type Researcher struct {
	Mc              MatchCrawler
	Sc              SummonerRankCrawler
	ChampionCrawler ChampionCrawler

	Comparer Comparer
}

type ResearchReq struct {
	CommonParam
	StartTime int64
	EndTime   int64
}

type ResearchRes struct {
	List []*ResearchResData
}

type ResearchResData struct {
	Key   string // championId+排序后的items
	Count int

	ChampionId int
	MatchInfo  MatchInfo
	Items      []Item
}

func (r *Researcher) Research(req ResearchReq) ResearchRes {
	res := ResearchRes{}
	summonIds := r.Sc.GetSummonerRank(GetSummonerRankReq{})
	for _, summonId := range summonIds {
		matchInfos, err := r.Mc.GetMatchList(GetMatchListReq{
			CommonParam: req.CommonParam,
			SummonerId:  summonId,
			StartTime:   req.StartTime,
			EndTime:     req.EndTime,
		})
		if err != nil {
			continue
		}

		for _, matchInfo := range matchInfos {
			// 缓存
			perf, err := r.ChampionCrawler.GetChampionItemPerformance(GetChampionItemPerformanceReq{})
			if err != nil {
				continue
			}

			diff := r.Comparer.Compare(matchInfo, perf)
			if diff {
				tmpKey := fmt.Sprintf("%d+%v", matchInfo.ChampionId, matchInfo.ItemList)
				for _, item := range res.List {
					if item.Key == tmpKey {
						item.Count++
						break
					}
				}
			}

		}
	}
	return res
}
