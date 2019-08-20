package main

import (
	"fmt"
	"texas-poker/match"
	"time"
)

var path = "./match/samples/seven_cards_with_ghost.result.json"

func main() {

	// 先读取数据
	mch := match.NewMatchs(path)
	startTime := time.Now()
	// 对比并且输出结果
	mch.CompareAndOutPutResults()
	endTime := time.Now()
	fmt.Printf("总数量：%d 条\n", len(mch.Matches))
	fmt.Printf("开始时间：%v 毫秒\n", startTime)
	fmt.Printf("结束时间：%v 毫秒\n", endTime)
	fmt.Printf("一共消耗：%.2f 毫秒\n", endTime.Sub(startTime).Seconds()*1000)
}
