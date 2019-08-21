package poker

// 手牌
type HandCard struct {
	cardType    uint64         // 牌的类型
	cardNum     int            // 牌的数量
	ghostNum    uint64         // 鬼牌数量
	score       uint64         // 手牌的分数
	cardStr     string         // 最原始的字符串牌型
	cards       [10]uint64     // 转换为数组切片的牌型，没有排序的
	faceSum     uint64         // 牌面的总和
	faceCounter [4]uint64      // 面牌次数计数器, 下标 6代表和，下标5记录逻辑与值
	colorTimes  [0x000f]uint64 // 记录花色出现的次数
}

func NewHandCard(cardStr string) *HandCard {
	hdc := &HandCard{}
	hdc.setCardsWithStr(cardStr)
	hdc.Analysis()
	return hdc
}

func (r *HandCard) Analysis()  {
	//// 1.检测是不是同花顺
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
	return
}

func (r *HandCard)GetCardType()uint64{
	return r.cardType
}
func (r *HandCard)GetScore()uint64{
	return r.score
}

// 设置手牌
func (r *HandCard) setCardsWithStr(cardStr string) {
	l := len(cardStr)
	r.ghostNum = 0
	r.faceSum = 0
	//r.cards = make([]uint64, 0, l)
	for i := 0; i < l; {
		ts := cardStr[i : i+2]
		cd := FaceAll[ts]
		i += 2
		r.faceSum += cd
		if cd&0x00f00000 > 0 {
			r.ghostNum++
		}
		val := cd & 0x0000ffff
		r.faceCounter[3] |= r.faceCounter[2] & val //记录出现过四次次以上的牌型，二进制 1 出现次数就是牌型的种类
		r.faceCounter[2] |= r.faceCounter[1] & val //记录出现过三次以上的牌型，二进制 1 出现次数就是牌型的种类
		r.faceCounter[1] |= r.faceCounter[0] & val //记录出现过两次以上的牌型，二进制 1 出现次数就是牌型的种类
		r.faceCounter[0] |= val                    //记录所有出现过的牌型, 二进制 1 出现次数就是牌型的种类

		r.colorTimes[(cd&0x000f0000) >> 16] |= val // 花色出现的次数

		r.cards[r.cardNum] = cd
		r.cardNum ++
	}
}

// 同花顺
func (r *HandCard) StraightFlush() bool {
	for i := uint64(0); i <= 4; i++ {
		cc := r.colorTimes[(0x0008)>>i]
		// 鬼牌最多只有两个，
		if CountBitNums(cc)+r.ghostNum >= 5 {
			r.cardType = TypeFlush
			if r.straight(r.faceSum, cc) {
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
	if r.faceCounter[3-r.ghostNum] > 0 {
		r.cardType = TypeFour
		fu := r.faceCounter[3-r.ghostNum]
		r.ToScore(GetFaceScore(fu) | GetHighNBits(fu^r.faceCounter[0], 15, 1))
		return true
	}
	return false
}

// 葫芦
func (r *HandCard) FullHose() bool {
	// 如果鬼为0个，那么必定是三条加一对，如果有一个鬼，必定是两对，如果是有两个鬼的话那么只要存在两对，就会变成四条
	if r.faceCounter[2-r.ghostNum] > 0 && CountBitNums(r.faceCounter[1]) > 1 {
		r.cardType = TypeFullHose
			highV := HighBitValue(r.faceCounter[2-r.ghostNum])
			two := ^(highV) & r.faceCounter[1]
			typeV := GetFaceScore(highV << 13 | two)
			r.ToScore(typeV)

		return true
	}
	return false
}

// 同花, 0 非同花， 大于0是同花，返回花色
func (r *HandCard) SameColor() bool {
	//colorNum := CountBitNums((r.faceColor & 0x000F0000) >> 8)
	//if colorNum == 1 && r.cardNum == 5 {
	//	r.cardType = TypeFlush
	//	// faceColor & 0x000F0000 花色
	//	r.ToScore(r.faceCounter[0])
	//	return true
	//}
	//
	for i := uint64(0); i <= 4; i++ {
		cc := r.colorTimes[0x0008>>i]
		// 鬼牌最多只有两个，
		if CountBitNums(cc)+r.ghostNum >= 5 {
			r.cardType = TypeFlush
			// 鬼牌可能充当的牌，如果不存在 AK 中的任意一个，鬼牌就可以充当，如果没有一个就充当QJT中的一个或两个
			canPull := (cc & TJQKA) ^ TJQKA
			canPull = GetHighNBits(canPull, 15, int(r.ghostNum)) + GetHighNBits(cc,15, int(5-r.ghostNum))
			r.ToScore(canPull)
			return true
		}
	}
	return false
}

// 顺子
func (r *HandCard) Straight() bool {
	return r.straight(r.faceSum, r.faceCounter[0])
}

// 三条
func (r *HandCard) ThreeSame() bool {
	// 有一个鬼的情况，必须要有一对存在，有两个鬼的话那么就可能变成四条了
	if r.faceCounter[2-r.ghostNum] > 0 {
		r.cardType = TypeThree
		one := r.faceCounter[2-r.ghostNum] ^ r.faceCounter[0]
		r.ToScore(GetFaceScore(r.faceCounter[2-r.ghostNum]) | GetHighNBits(one, 14, 2))
		return true
	}
	return false
}

// straight 判断是不是顺子, 包括有鬼牌的情况
func (r *HandCard) straight(sum uint64, card uint64) bool {
	cModel := uint64(TJQKA)
	if r.ghostNum == 0 {
		card = card & 0x0000ffff
		for cModel >= 31 {
			if cModel&card == cModel {
				r.cardType = TypeStraight
				r.ToScore(cModel)
				return true
			}
			cModel = cModel >> 1
		}
	}
	for cModel >= 31 {
		if oneNum := CountBitNums(cModel & card); oneNum >= 5-r.ghostNum {
			r.cardType = TypeStraight
			r.ToScore(cModel)
			return true
		}
		cModel = cModel >> 1
	}
	// 最小的顺子
	if oneNum := CountBitNums(card & A2345); oneNum >= 5-r.ghostNum {
		r.cardType = TypeStraight
		r.ToScore(A2345&(^uint64(0x00002000))|1)
		return true
	}

	return false
}

// 两对
func (r *HandCard) TwoPairs() bool {
	// 如果存在鬼牌，且存在一对以上的话必定会变成三条，所有不考虑鬼牌情况
	if CountBitNums(r.faceCounter[1]) > 1 {
		r.cardType = TypeTwoPair
		highV := HighBitValue(r.faceCounter[1])
		two := HighBitValue(^(highV) & r.faceCounter[1])
		typeV := GetFaceScore(highV|two)
		//base  := GetHighNBits(r.faceCounter[0]^highV^two, 16, 1)
		base  := HighBitValue(r.faceCounter[0]^highV^two)
		r.ToScore(typeV | base)
		return true
	}
	return false
}

// 一对
func (r *HandCard) OnePairs() bool {
	//如果存在一张鬼牌那必定会有一对，
	if r.ghostNum > 0 || r.faceCounter[1] > 0 {
		r.cardType = TypeOnePair
		one := HighBitValue(r.faceCounter[1-r.ghostNum])
		r.ToScore(GetHighNBits(one^r.faceCounter[0], 14, 3) | GetFaceScore(one))
		return true
	}
	return false
}

// HighCard 高牌
func (r *HandCard) HighCard() bool {
	r.cardType = TypeHighCard
	r.ToScore(GetHighNBits(r.faceCounter[0], 16, 5))
	return true
}

func (r *HandCard) ToScore(base uint64) {
	r.score = r.cardType  | base
}
