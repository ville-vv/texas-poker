package poker

var FaceAll = map[string]uint32{
	"2d": 0x00010000 + 1<<1,
	"2c": 0x00020000 + 1<<1,
	"2h": 0x00040000 + 1<<1,
	"2s": 0x00080000 + 1<<1,
	"3d": 0x00010000 + 1<<2,
	"3c": 0x00020000 + 1<<2,
	"3h": 0x00040000 + 1<<2,
	"3s": 0x00080000 + 1<<2,
	"4d": 0x00010000 + 1<<3,
	"4c": 0x00020000 + 1<<3,
	"4h": 0x00040000 + 1<<3,
	"4s": 0x00080000 + 1<<3,
	"5d": 0x00010000 + 1<<4,
	"5c": 0x00020000 + 1<<4,
	"5h": 0x00040000 + 1<<4,
	"5s": 0x00080000 + 1<<4,
	"6d": 0x00010000 + 1<<5,
	"6c": 0x00020000 + 1<<5,
	"6h": 0x00040000 + 1<<5,
	"6s": 0x00080000 + 1<<5,
	"7d": 0x00010000 + 1<<6,
	"7c": 0x00020000 + 1<<6,
	"7h": 0x00040000 + 1<<6,
	"7s": 0x00080000 + 1<<6,
	"8d": 0x00010000 + 1<<7,
	"8c": 0x00020000 + 1<<7,
	"8h": 0x00040000 + 1<<7,
	"8s": 0x00080000 + 1<<7,
	"9d": 0x00010000 + 1<<8,
	"9c": 0x00020000 + 1<<8,
	"9h": 0x00040000 + 1<<8,
	"9s": 0x00080000 + 1<<8,
	"Td": 0x00010000 + 1<<9,
	"Tc": 0x00020000 + 1<<9,
	"Th": 0x00040000 + 1<<9,
	"Ts": 0x00080000 + 1<<9,
	"Jd": 0x00010000 + 1<<10,
	"Jc": 0x00020000 + 1<<10,
	"Jh": 0x00040000 + 1<<10,
	"Js": 0x00080000 + 1<<10,
	"Qd": 0x00010000 + 1<<11,
	"Qc": 0x00020000 + 1<<11,
	"Qh": 0x00040000 + 1<<11,
	"Qs": 0x00080000 + 1<<11,
	"Kd": 0x00010000 + 1<<12,
	"Kc": 0x00020000 + 1<<12,
	"Kh": 0x00040000 + 1<<12,
	"Ks": 0x00080000 + 1<<12,
	"Ad": 0x00010000 + 1<<13,
	"Ac": 0x00020000 + 1<<13,
	"Ah": 0x00040000 + 1<<13,
	"As": 0x00080000 + 1<<13,
	"Xx": 0x00100000,
	"Xn": 0x00200000,
}