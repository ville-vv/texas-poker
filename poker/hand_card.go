package poker

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
	faceSum     uint32            // 牌面的总和
	faceColor   uint32            // 牌所有花色记录
	FaceCounter [7]uint32         // 面牌次数计数器, 下标 6代表和，下标5记录逻辑与值
	ColorTimes  map[uint32]uint32 // 记录花色出现的次数
	score       uint32            // 手牌的分数
}

func NewHandCard(cardStr string) *HandCard {
	hdc := &HandCard{}
	hdc.SetCardsWithStr(cardStr)
	return hdc
}

func (r *HandCard) GetMaxCardScore() uint32 {
	// 1.检测是不是同花顺
	if r.StraightFlush() {
	} else if r.FourSame() {
	} else if r.FullHose() {
		// 3.检测是不是葫芦
	} else if r.SameColor() {
		// 4.检测是不是同花
	} else if r.Straight() {
		// 5.检测是不是顺子
	} else if r.ThreeSame() {
		// 6.检测是不是三条
	} else if r.TwoPairs() {
		// 7.检测是不是两对
	} else if r.OnePairs() {
		// 7.检测是不是一对
	} else {
		// 剩下就是高牌
		r.HighCard()
	}
	return r.score
}

// 设置手牌
func (r *HandCard) SetCards(c []uint32) {
	l := len(c)
	var sum uint32
	var color uint32
	r.faceSum = 0
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
	r.faceColor = color
	r.faceSum = sum
	r.cards = c[:]
	r.cardNum = l
}

// 设置手牌
func (r *HandCard) SetCardsWithStr(cardStr string) {
	l := len(cardStr)
	r.ColorTimes = make(map[uint32]uint32)
	r.faceColor = 0
	r.GhostNum = 0
	r.faceSum = 0
	r.FaceCounter = [7]uint32{0}
	r.cards = make([]uint32, 0, l)
	for i := 0; i < l; {
		ts := cardStr[i : i+2]
		cd := FaceAll[ts]
		i += 2
		r.faceSum += cd
		r.faceColor |= cd
		if cd&0x00f00000 > 0 {
			r.GhostNum++
		}

		val := cd & 0x0000ffff
		r.FaceCounter[3] |= r.FaceCounter[2] & val //记录出现过四次次以上的牌型，二进制 1 出现次数就是牌型的种类
		r.FaceCounter[2] |= r.FaceCounter[1] & val //记录出现过三次以上的牌型，二进制 1 出现次数就是牌型的种类
		r.FaceCounter[1] |= r.FaceCounter[0] & val //记录出现过两次以上的牌型，二进制 1 出现次数就是牌型的种类
		r.FaceCounter[0] |= val                    //记录所有出现过的牌型, 二进制 1 出现次数就是牌型的种类

		r.ColorTimes[cd&0x000f0000] |= val // 花色出现的次数

		r.cards = append(r.cards, cd)
	}
	r.cardNum = len(r.cards)
}

// 皇家同花顺
func (r *HandCard) KingStraightFlush() bool {
	switch r.faceSum {
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
	for i := uint32(0); i <= 4; i++ {
		// 鬼牌最多只有两个，
		if CountBitNums(r.ColorTimes[0x00080000>>i])+r.GhostNum >= 5 {
			if r.straight(r.faceSum, r.ColorTimes[0x00080000>>i]) {
				r.cardType = TypeStraightFlush
				r.ToScore(r.score | 0x00ffffff)
				return true
			}

		}
	}
	return false
}

// 四条
func (r *HandCard) FourSame() bool {
	if r.FaceCounter[3-r.GhostNum] > 0 {
		r.cardType = TypeFour
		r.ToScore(GetFaceScore(r.FaceCounter[3-r.GhostNum]) | r.FaceCounter[0])
		return true
	}
	return false
}

// 葫芦
func (r *HandCard) FullHose() bool {
	// 如果鬼为0个，那么必定是三条加一对，如果有一个鬼，必定是两对，如果是有两个鬼的话那么只要存在两对，就会变成四条
	if r.FaceCounter[2-r.GhostNum] > 0 && CountBitNums(r.FaceCounter[1]) > 1 {
		r.cardType = TypeFullHose
		highV := HighBitValue(r.FaceCounter[2-r.GhostNum])
		two := ^(highV) & r.FaceCounter[1-r.GhostNum]
		r.ToScore(GetFaceScore(highV) + GetFaceScore(two) | GetHighNBits(r.FaceCounter[0]^highV^two, 14, 3))
		return true
	}
	return false
}

// 同花, 0 非同花， 大于0是同花，返回花色
func (r *HandCard) SameColor() bool {
	colorNum := CountBitNums((r.faceColor & 0x000F0000) >> 8)
	if colorNum == 1 && r.cardNum == 5 {
		r.cardType = TypeFlush
		// faceColor & 0x000F0000 花色
		r.ToScore(r.FaceCounter[0])
		return true
	}
	for i := uint32(0); i <= 4; i++ {
		// 鬼牌最多只有两个，
		if CountBitNums(r.ColorTimes[0x00080000>>i])+r.GhostNum >= 5 {
			// 0x00080000 >> i
			r.cardType = TypeFlush
			// 鬼牌可能充当的牌，如果不存在 AK 中的任意一个，鬼牌就可以充当，如果没有一个就充当QJT中的一个或两个
			canPull := (r.FaceCounter[0] & TJQKA) ^ TJQKA
			canPull = GetHighNBits(canPull, 15, int(r.GhostNum)) | r.FaceCounter[0]
			r.ToScore(canPull)
			return true
		}
	}
	return false
}

// 顺子
func (r *HandCard) Straight() bool {
	return r.straight(r.faceSum, r.FaceCounter[0])
}

// 三条
func (r *HandCard) ThreeSame() bool {
	// 有一个鬼的情况，必须要有一对存在，有两个鬼的话那么就可能变成四条了
	if r.FaceCounter[2-r.GhostNum] > 0 {
		r.cardType = TypeThree
		one := r.FaceCounter[2-r.GhostNum] ^ r.FaceCounter[0]
		r.ToScore(GetFaceScore(r.FaceCounter[2-r.GhostNum]) | GetHighNBits(one, 16, 2))
		return true
	}
	return false
}

// straight 判断是不是顺子, 包括有鬼牌的情况
func (r *HandCard) straight(sum uint32, card uint32) bool {
	cModel := uint32(TJQKA)
	if r.GhostNum == 0 {
		sum = sum & 0x0000ffff
		for cModel >= 31 {
			if cModel&sum == cModel {
				r.cardType = TypeStraight
				r.ToScore(cModel)
				return true
			}
			cModel = cModel >> 1
		}
	}
	for cModel >= 31 {
		if oneNum := CountBitNums(cModel & card); oneNum >= 5-r.GhostNum {
			r.cardType = TypeStraight
			r.ToScore(cModel)
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
		highV := HighBitValue(r.FaceCounter[1])
		two := ^(highV) & r.FaceCounter[1]
		r.ToScore(GetFaceScore(highV) + GetFaceScore(two) | GetHighNBits(r.FaceCounter[0]^highV^two, 16, 1))
		return true
	}
	return false
}

// 一对
func (r *HandCard) OnePairs() bool {
	//如果存在一张鬼牌那必定会有一对，
	if r.GhostNum > 0 || r.FaceCounter[1] > 0 {
		r.cardType = TypeOnePair
		one := r.FaceCounter[1] ^ r.FaceCounter[0]
		r.ToScore(GetHighNBits(one, 16, 3) + GetFaceScore(r.FaceCounter[1]))
		return true
	}
	return false
}

// HighCard 高牌
func (r *HandCard) HighCard() bool {
	r.cardType = TypeHighCard
	r.ToScore(GetHighNBits(r.FaceCounter[0], 16, 5))
	return true
}

func (r *HandCard) ToScore(base uint32) {
	r.score = (0x01000000 * r.cardType) | base
}
