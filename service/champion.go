package service

// 装备表现
type ItemPerformance struct {
	Item     Item
	ShowRate int
	WinRate  int
}

type ChampionItemPerformance struct {
	One  ItemPerformance
	Two  []ItemPerformance
	Tree []ItemPerformance
	Four []ItemPerformance
}

type GetChampionItemPerformanceReq struct {
	CommonParam
	ChampionId int
	Lane       string
}

// 获取英雄常规出装的接口
type ChampionCrawler interface {
	GetChampionItemPerformance(req GetChampionItemPerformanceReq) (ChampionItemPerformance, error)
}
