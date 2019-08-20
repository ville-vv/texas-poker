package match

import (
	"fmt"
	"github.com/json-iterator/go"
	"io/ioutil"
	"texas-poker/poker"
)

type Match struct {
	Alice string `json:"alice"`
	Bob string `json:"bob"`
	Result  int    `json:"result"`
}


type Matches struct {
	Matches []*Match `json:"matches"`
}

func NewMatchs(fPath string)*Matches{
	m := &Matches{}
	ReadMatchFromJson(fPath,m)
	return m
}

func (matches *Matches) OutPutResults() {
	cnter := 0
	for _, v := range matches.Matches {
		res := Compare(v.Alice, v.Bob)
		if res != v.Result {
			fmt.Printf("%s, %s , %d, %d\n", v.Alice, v.Bob, res, v.Result)
			cnter++
		}
	}
	if cnter > 0{
		fmt.Printf("不正确结果为：%d 条\n", cnter)
	}
}

func ReadMatchFromJson(path string, m *Matches ) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic("panic: " + err.Error())
	}
	if err := jsoniter.Unmarshal(file, m); err != nil {
		panic("panic: " + err.Error())
	}
	return
}

func Compare(strA string, strB string) int {
	card1 := poker.NewHandCard(strA)
	card2 := poker.NewHandCard(strB)
	return getResult(card1.GetScore(), card2.GetScore())
}
// 获取获胜者编号
func getResult(a, b uint64) int {
	switch {
	case a > b:
		return 1
	case a < b:
		return 2
	}
	return 0
}
