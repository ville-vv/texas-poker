package poker

import "fmt"

const (
	CompColorFlag = 0x000F0000
)

const (
	SpadesKingStraightFlush   uint32 = 0x53E00  // 0x00010200+0x00010400+0x00010800+0x00011000+0x00012000
	HeartsKingStraightFlush   uint32 = 0xA3E00  // 0x00020200+0x00020400+0x00020800+0x00021000+0x00022000
	DiamondsKingStraightFlush uint32 = 0x143E00 // 0x00040200+0x00040400+0x00040800+0x00041000+0x00042000
	SlubsKingStraightFlush    uint32 = 0x283E00 // 0x00080200+0x00080400+0x00080800+0x00081000+0x00082000
	TJQKA                     uint32 = 0x283E00 & 0x0000ffff
)

// 手牌
type HandCard struct {
	cardType    uint32            // 牌的类型
	cardNum     int               // 牌的数量
	GhostNum    uint32            // 鬼牌数量
	CardStr     string            // 最原始的字符串牌型
	cards       []uint32          // 转换为数组切片的牌型，没有排序的
	FaceSum     uint32            // 牌面的总和
	FaceColor   uint32            // 牌所有花色记录
	FaceCounter [7]uint32         // 面牌次数计数器, 下标 6代表和，下标5记录逻辑与值
	ColorTimes  map[uint32]uint32 // 记录花色出现的次数
	score       uint32            // 手牌的分数
}

func NewRule() *HandCard {
	return &HandCard{}
}

// 设置手牌
func (r *HandCard) SetCards(c []uint32) {
	l := len(c)
	var sum uint32
	var color uint32
	r.FaceSum = 0
	r.FaceCounter = [7]uint32{0}
	for i := 0; i < l; i++ {
		sum += c[i]
		color |= c[i]

		val := c[i] & 0x0000ffff
		r.FaceCounter[3] |= r.FaceCounter[2] & val //记录出现过四次次以上的牌型，二进制 1 出现次数就是牌型的种类
		r.FaceCounter[2] |= r.FaceCounter[1] & val //记录出现过三次以上的牌型，二进制 1 出现次数就是牌型的种类
		r.FaceCounter[1] |= r.FaceCounter[0] & val //记录出现过两次以上的牌型，二进制 1 出现次数就是牌型的种类
		r.FaceCounter[0] |= val                    //记录所有出现过的牌型, 二进制 1 出现次数就是牌型的种类
	}
	r.FaceColor = color
	r.FaceSum = sum
	r.cards = c[:]
	r.cardNum = l
}

// 设置手牌
func (r *HandCard) SetCardsWithStr(cardStr string) {
	l := len(cardStr)
	r.ColorTimes = make(map[uint32]uint32)
	r.FaceColor = 0
	r.GhostNum = 0
	r.FaceSum = 0
	r.FaceCounter = [7]uint32{0}
	r.cards = make([]uint32, 0, l)
	for i := 0; i < l; {
		ts := cardStr[i : i+2]
		cd := FaceAll[ts]
		i += 2
		r.FaceSum += cd
		r.FaceColor |= cd
		if cd&0x00f00000 > 0 {
			r.GhostNum++
		}

		val := cd & 0x0000ffff
		r.FaceCounter[3] |= r.FaceCounter[2] & val //记录出现过四次次以上的牌型，二进制 1 出现次数就是牌型的种类
		r.FaceCounter[2] |= r.FaceCounter[1] & val //记录出现过三次以上的牌型，二进制 1 出现次数就是牌型的种类
		r.FaceCounter[1] |= r.FaceCounter[0] & val //记录出现过两次以上的牌型，二进制 1 出现次数就是牌型的种类
		r.FaceCounter[0] |= val                    //记录所有出现过的牌型, 二进制 1 出现次数就是牌型的种类

		r.ColorTimes[cd&0x000f0000] += 1 // 花色出现的次数

		r.cards = append(r.cards, cd)
	}
	r.cardNum = len(r.cards)
}

// 皇家同花顺
func (r *HandCard) KingStraightFlush() bool {
	switch r.FaceSum {
	case SpadesKingStraightFlush:
		return true
	case HeartsKingStraightFlush:
		return true
	case DiamondsKingStraightFlush:
		return true
	case SlubsKingStraightFlush:
		return true
	}

	return false
}

// 同花顺
func (r *HandCard) StraightFlush() bool {
	if r.SameColor() == 0 {
		return false
	}
	return r.straight(r.FaceSum)
}

// 四条
func (r *HandCard) FourSame() bool {
	if r.FaceCounter[3-r.GhostNum] > 0 {
		return true
	}
	return false
}

// 葫芦
func (r *HandCard) FullHose() bool {
	// 如果鬼为0个，那么必定是三条加一对，如果有一个鬼，必定是两对，如果是有两个鬼的话那么只要存在两对，就会变成四条
	if r.FaceCounter[2-r.GhostNum] > 0 && CountBitNums(r.FaceCounter[1]) > 1 {
		return true
	}
	return false
}

// 同花, 0 非同花， 大于0是同花，返回花色
func (r *HandCard) SameColor() uint32 {
	colorNum := CountBitNums((r.FaceColor & 0x000F0000) >> 8)
	if colorNum == 1 && r.cardNum == 5 {
		return r.FaceColor & 0x000F0000
	}
	fmt.Println(r.ColorTimes)
	for i := uint32(0); i < 4; i++ {
		if r.ColorTimes[0x00080000>>i]+r.GhostNum == 5 {
			return 0x00080000 >> i
		}
	}
	return 0
}

// 顺子
func (r *HandCard) Straight() (bool, uint32) {
	d := r.FaceCounter[0] & 0x0000ffff
	max := HighBitValue(d)
	return r.straight(r.FaceSum), max
}

// 三条
func (r *HandCard) ThreeSame() bool {
	// 有一个鬼的情况，必须要有一对存在，有两个鬼的话那么就可能变成四条了
	if r.FaceCounter[2-r.GhostNum] > 0 {
		return true
	}
	return false
}

// straight 判断是不是顺子, 包括有鬼牌的情况
func (r *HandCard) straight(sum uint32) bool {
	cModel := uint32(TJQKA)
	if r.GhostNum == 0 {
		sum = sum & 0x0000ffff
		for cModel >= 31 {
			if cModel&sum == cModel {
				r.cardType = TypeStraight
				//r.ToScore()
				return true
			}
			cModel = cModel >> 1
		}
	}
	for cModel >= 31 {
		if oneNum := CountBitNums(cModel & r.FaceCounter[0]); oneNum >= 5-r.GhostNum {
			return true
		}
		cModel = cModel >> 1
	}

	return false
}

// 两对
func (r *HandCard) TwoPairs() bool {
	// 如果存在鬼牌，且存在一对以上的话必定会变成三条，所有不考虑鬼牌情况
	if CountBitNums(r.FaceCounter[1]) > 1 {
		r.cardType = TypeTwoPair
		r.ToScore(r.FaceCounter[1])
		return true
	}
	return false
}

// 一对
func (r *HandCard) OnePairs() bool {
	//如果存在一张鬼牌那必定会有一对，
	if r.GhostNum > 0 || r.FaceCounter[1] > 0 {
		r.cardType = TypeOnePair
		r.ToScore(GetHighNBits(r.FaceSum, 16, 5))
		return true
	}
	return false
}

func (r *HandCard) HighCard() bool {
	if (r.FaceCounter[0] & 0x0000ffff) == uint32(r.cardNum) {
		r.cardType = TypeHighCard
		r.ToScore(GetHighNBits(r.FaceSum, 16, 5))
		return true
	}
	return false
}

func (r *HandCard) ToScore(v uint32) {
	r.score = (0x00010000 << r.cardType) + v
}
