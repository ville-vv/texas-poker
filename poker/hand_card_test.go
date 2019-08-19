package poker

import (
	"fmt"
	"testing"
)

func TestRule_SameColor(t *testing.T) {
	cardStrings := []string{"2s3s4c5cXnXnXn", "5h7h6h8h9h", "TcJcQcKcAc", "TdJdQdKdAd"}
	for _, v := range cardStrings {
		rule := HandCard{}
		rule.SetCardsWithStr(v)
		if rule.SameColor(){
			t.Logf("%s 是同花", v)
		}
		//switch  {
		//case 0x00010000:
		//	t.Logf("%s 是方块同花", v)
		//case 0x00020000:
		//	t.Logf("%s 是梅花同花", v)
		//case 0x00040000:
		//	t.Logf("%s 是红心同花", v)
		//case 0x00080000:
		//	t.Logf("%s 是黑桃同花", v)
		//default:
		//	t.Logf("%s 不是同花", v)
		//}
	}
}

func TestRule_Straight(t *testing.T) {
	cardStrings := []string{"9hXnJhQhXn5c7c", "4h7h6h8h9c", "5hTh6h8h9c", "8h9hTh7hJc"}
	//cardStrings = []string{"8h9hTh7hJc"}
	for _, v := range cardStrings {
		rule := HandCard{}
		rule.SetCardsWithStr(v)
			if rule.Straight(){
				t.Logf("%s 是顺子", v)
			}else{
				t.Logf("%s 不是顺子", v)
			}

	}
}

func TestRule_KingStraightFlush(t *testing.T) {
	rule := HandCard{}
	cardStrings := []string{"2s3s4s6s5s", "ThJhQhKhAh", "TcJcQcKcAc", "TdJdQdKdAd"}
	for _, v := range cardStrings {
		rule.SetCards(StringToCards(v))
		t.Logf("%s 是不是皇家同花顺: %v", v, rule.KingStraightFlush())
	}
}

func TestRule_FourSame(t *testing.T) {

	cardStrings := []string{"2s2d2c6sXx", "ThJhQhKhAhXxXxXx", "TcJcQcKcAc", "TdJdQdKdAd"}
	for _, v := range cardStrings {
		rule := &HandCard{}
		rule.SetCardsWithStr(v)
		if rule.FourSame() {
			fmt.Printf("%v, 是四条\n", v)
		} else {
			fmt.Printf("%v, 不是四条\n", v)
		}
	}

}

func TestHandCard_FullHose(t *testing.T) {
	cardStrings := []string{"6h6h5h5h3hXx9h","6h6h6h5h5h7h9h","6h6h4h5h5h7h9h"}
	for _, v := range cardStrings {
		rule := NewHandCard(v)
		t.Logf("%s 是否葫芦：%v", v, rule.FullHose())
	}
}

func TestRule_ThreeSame(t *testing.T) {
	cardStrings := []string{"6h6h4h5h3hXx9hXn"}
	for _, v := range cardStrings {
		rule := NewHandCard(v)
		t.Logf("%s 是否三个：%v", v, rule.ThreeSame())
	}
}

func TestHandCard_TwoPairs(t *testing.T) {
	cardStrings := []string{"7h6h4h4h3h3c9h","6h6h6h5h5h7h9h","6h6h4h5h5h7h9h"}
	for _, v := range cardStrings {
		rule := NewHandCard(v)
		fmt.Printf("%s 是否两队：%v\n", v, rule.TwoPairs())
	}
}

func TestHandCard_OnePairs(t *testing.T) {
	cardStrings := []string{"7h6h4h4h3h3c9h","6h6h6h5h5h7h9h","6h6h4h5h5h7h9h"}
	for _, v := range cardStrings {
		rule := NewHandCard(v)
		fmt.Printf("%s 是否两队：%v\n", v, rule.OnePairs())
	}
}