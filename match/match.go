package match

import (
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

func (m *Matches) CompareAndOutPutResults() {
	counter := 0
	for _, v := range m.Matches {
		card1 := poker.NewHandCard(v.Alice)
		card2 := poker.NewHandCard(v.Bob)
		if res := m.getResult(card1.GetScore(), card2.GetScore()); res!= v.Result {
			//fmt.Printf("%s, %s , %d, %d\n", v.Alice, v.Bob, res, v.Result)
			counter++
		}
	}
	//if counter > 0{
	//	fmt.Printf("不正确结果为：%d 条\n", counter)
	//}
}

func (m *Matches)getResult(a, b uint64) int {
	switch {
	case a > b:
		return 1
	case a < b:
		return 2
	}
	return 0
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
