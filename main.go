package main

import (
	"fmt"
	"texas-poker/match"
	"time"
)

var path = map[string]string{
	"五张没赖子":"./match/samples/match.json",
	"五张加赖子":"./match/samples/five_cards_with_ghost.result.json",
	"七张没赖子":"./match/samples/seven_cards.json",
	"七张加赖子":"./match/samples/seven_cards_with_ghost.result.json",
}

func main() {
	for k ,v := range path{
		// 先读取数据
		mch := match.NewMatchs(v)
		startTime := time.Now()
		// 对比并且输出结果
		mch.CompareAndOutPutResults()
		endTime := time.Now()
		fmt.Println("============================================================")
		fmt.Printf("类型：%s 条\n", k)
		fmt.Printf("总数量：%d 条\n", len(mch.Matches))
		fmt.Printf("开始时间：%v 毫秒\n", startTime)
		fmt.Printf("结束时间：%v 毫秒\n", endTime)
		fmt.Printf("一共消耗：%.2f 毫秒\n", endTime.Sub(startTime).Seconds()*1000)
		fmt.Println("============================================================")
	}

}
