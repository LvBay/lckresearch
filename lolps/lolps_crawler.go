package lolps

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
)

type SummonerData struct {
	Summoners []string `json:"summoners"`
}

func GetTop10Summoners() error {
	// 创建带取消的上下文
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// 设置超时时间
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var summoners []string

	// 执行爬虫任务
	err := chromedp.Run(ctx,
		// 访问目标网页
		chromedp.Navigate(`https://lol.ps/ranking?lane=-1&page=1&region=kr`),
		// 等待表格加载
		chromedp.WaitVisible(`table.table`),
		// 等待确保数据加载完成
		chromedp.Sleep(2*time.Second),
		// 提取前10名召唤师名称
		chromedp.Evaluate(`
			Array.from(document.querySelectorAll('table.table tbody tr'))
				.slice(0, 10)
				.map(row => row.querySelector('td:nth-child(2) a').textContent)
		`, &summoners),
	)

	if err != nil {
		return fmt.Errorf("爬取数据失败: %v", err)
	}

	// 准备数据
	data := SummonerData{
		Summoners: summoners,
	}

	// 将结果保存到JSON文件
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("JSON编码失败: %v", err)
	}

	err = os.WriteFile("ids.json", jsonData, 0644)
	if err != nil {
		return fmt.Errorf("保存文件失败: %v", err)
	}

	fmt.Printf("爬取完成，共获取到 %d 个召唤师ID\n", len(summoners))
	return nil
}
