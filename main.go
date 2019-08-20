package main

import (
	"fmt"
	"texas-poker/match"
	"time"
)

var path = "./match/samples/seven_cards_with_ghost.result.json"

func main() {
	mch := match.NewMatchs(path)
	startTime := time.Now()
	mch.OutPutResults()
	endTime := time.Now()
	fmt.Printf("开始时间：%v 毫秒\n", startTime)
	fmt.Printf("结束时间：%v 毫秒\n", endTime)
	fmt.Printf("一共消耗：%.2f 毫秒\n", endTime.Sub(startTime).Seconds()*1000)
}
