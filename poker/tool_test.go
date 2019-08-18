package poker

import "testing"

func TestHighBitValue(t *testing.T) {
	if rs := HighBitValue(0x5); rs != 0x4 {
		t.Errorf("获取最高位结果不正确:%d", rs)
	}
	if rs := HighBitValue(0x9); rs != 0x8 {
		t.Errorf("获取最高位结果不正确:%d", rs)
	}
	if rs := HighBitValue(0x12); rs != 0x10 {
		t.Errorf("获取最高位结果不正确:%d", rs)
	}
	if rs := HighBitValue(0x22); rs != 0x20 {
		t.Errorf("获取最高位结果不正确:%d", rs)
	}
}

func TestGetHighNBits(t *testing.T) {
	vl := uint32(0xAB00)
	t.Log(GetHighNBits(vl, 16, 5))
}
