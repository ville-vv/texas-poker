package poker

import (
	"fmt"
	"github.com/json-iterator/go"
	"io/ioutil"
	"time"
)

type Matches struct {
	Matches []*Match `json:"matches"`
}

type Match struct {
	PlayerA string `json:"alice"`
	PlayerB string `json:"bob"`
	Result  int    `json:"result"`
}

// 获取牌组（必然获取，否则抛出异常）
func MustGetMatchesFromMatchSamples(path string) *Matches {
	var file []byte
	var err error

	if file, err = ioutil.ReadFile(path); err != nil {
		panic("panic: " + err.Error())
	}

	matches := Matches{}
	if err := jsoniter.Unmarshal(file, &matches); err != nil {
		panic("panic: " + err.Error())
	}
	return &matches
}

// 打印牌组比较结果
func (matches *Matches) PrintCompareResult() {
	beginTime := time.Now()
	cnter := 0
	for _, v := range matches.Matches {
		res := Compare(v.PlayerA, v.PlayerB)
		if res != v.Result {
			fmt.Printf("%s, %s , %d, %d\n", v.PlayerA, v.PlayerB, res, v.Result)
			cnter++
		}
	}
	finishTime := time.Now()
	fmt.Printf("共耗时：%.2f 毫秒\n", finishTime.Sub(beginTime).Seconds()*1000)
	fmt.Printf("合计：%d 条\n", len(matches.Matches))
	fmt.Printf("不正确结果为：%d 条\n", cnter)
}

// 比较两张手牌、支持任意数量手牌及任意数量赖子
func Compare(strA string, strB string) int {
	playerA := NewHandCard(strA).GetMaxCardScore()
	playerB := NewHandCard(strB).GetMaxCardScore()
	// 比较最大牌型
	if winner := getWinner(playerA, playerB); winner != 0 {
		return winner
	}
	return 0
}

// 获取获胜者编号
func getWinner(a, b uint32) int {
	switch {
	case a > b:
		return 1
	case a < b:
		return 2
	}
	return 0
}
