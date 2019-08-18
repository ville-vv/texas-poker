package poker

const (
	TypeRoyalFlush    uint32 = 10 // 皇家同花顺：同一花色的最大顺子。（最大牌：A-K-Q-J-10）
	TypeStraightFlush uint32 = 9  // 同花顺：同一花色的顺子。（最大牌：K-Q-J-10-9 最小牌：A-2-3-4-5）
	TypeFour          uint32 = 8  // 四条：四同张加单张。（最大牌：A-A-A-A-K 最小牌：2-2-2-2-3）
	TypeFullHose      uint32 = 7  // 葫芦（豪斯）：三同张加对子。（最大牌：A-A-A-K-K 最小牌：2-2-2-3-3）
	TypeFlush         uint32 = 6  // 同花：同一花色。（最大牌：A-K-Q-J-9 最小牌：2-3-4-5-7）
	TypeStraight      uint32 = 5  // 顺子：花色不一样的顺子。（最大牌：A-K-Q-J-10 最小牌：A-2-3-4-5）
	TypeThree         uint32 = 4  // 三条：三同张加两单张。（最大牌：A-A-A-K-Q 最小牌：2-2-2-3-4）
	TypeTwoPair       uint32 = 3  // 两对：（最大牌：A-A-K-K-Q 最小牌：2-2-3-3-4）
	TypeOnePair       uint32 = 2  // 一对：（最大牌：A-A-K-Q-J 最小牌：2-2-3-4-5）
	TypeHighCard      uint32 = 1  // 高牌：（最大牌：A-K-Q-J-9 最小牌：2-3-4-5-7）
)
