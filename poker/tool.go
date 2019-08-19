package poker

func StringToCards(cs string) []uint32 {
	l := len(cs)
	cards := make([]uint32, 0, l)
	for i := 0; i < l; {
		ts := cs[i : i+2]
		cards = append(cards, FaceAll[ts])
		i = i + 2
	}
	return cards
}

// 获取一个数的二进制 1 的数量位数
func CountBitNums(a uint32) uint32 {
	a = (0x55555555 & a) + ((a >> 1) & 0x55555555)
	a = (0x33333333 & a) + ((a >> 2) & 0x33333333)
	a = (0x0f0f0f0f & a) + ((a >> 4) & 0x0f0f0f0f)
	a = (0x00ff00ff & a) + ((a >> 8) & 0x00ff00ff)
	return a
}

// 获取一个数最低位的值
func LeastBitValue(n uint32) uint32 {
	res := n
	n = n & (n - 1)
	return ^n & res
}

// 获取一个数最高位的值
func HighBitValue(x uint32) uint32 { //0010 1100 0000 0000 0000 0000 0000 0000 0000 0001
	x = x | (x >> 1)  //0011 1110 0000 0000 0000 0000 0000 0000 0000 0000
	x = x | (x >> 2)  //0011 1111 1000 0000 0000 0000 0000 0000 0000 0000
	x = x | (x >> 4)  //0011 1111 1111 1000 0000 0000 0000 0000 0000 0000
	x = x | (x >> 8)  //0011 1111 1111 1111 1111 1000 0000 0000 0000 0000
	x = x | (x >> 16) //0011 1111 1111 1111 1111 1111 1111 1111 1111 1111
	x = x | (x >> 32)
	return (x + 1) >> 1 //0100 0000 0000 0000 0000 0000 0000 0000 0000 0000
}

// 获取m到n位最大bit组合的值
func GetHighNBits(v uint32, m, n int) uint32 {
	c := uint32(0x00008000)
	w := uint32(0)
	for i := 0; i < m && n > 0; i++ {
		if v&c > 0 {
			w |= c
			n--
		}
		c = c >> 1
	}
	return w
}

func GetFaceScore(v uint32) uint32 {
	v = HighBitValue(v)
	var cnt uint32
	for v > 0 {
		if (v & 1) > 0 {
			return (1 << 13) * cnt
		}
		v = v >> 1
		cnt++
	}
	return 0
}
