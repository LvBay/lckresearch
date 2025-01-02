package lolps

type PageInfo struct {
	Page  int `json:"page"`
	Start int `json:"start"`
	End   int `json:"end"`
}

func getPageInfo(start, end, pageSize int) []PageInfo {
	var result []PageInfo

	// 计算起始页和结束页
	startPage := (start-1)/pageSize + 1
	endPage := (end-1)/pageSize + 1

	for page := startPage; page <= endPage; page++ {
		// 计算当前页的全局起始和结束范围
		pageStart := (page-1)*pageSize + 1
		pageEnd := page * pageSize

		// 修正范围
		actualStart := max(start, pageStart)
		actualEnd := min(end, pageEnd)

		// 修改：end值减1，使其符合从0开始计数的逻辑
		pageRelativeStart := actualStart - pageStart + 1
		pageRelativeEnd := actualEnd - pageStart // 移除了 +1

		result = append(result, PageInfo{
			Page:  page,
			Start: pageRelativeStart,
			End:   pageRelativeEnd,
		})
	}

	return result
}
