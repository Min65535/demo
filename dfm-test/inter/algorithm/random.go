package algorithm

type MersenneTwister struct {
	index int
	MT    [624]int // 624 * 32 - 31 = 19937
}

func NewMersenneTwister(seed int) MersenneTwister {
	mt := MersenneTwister{}
	mt.index = 0
	mt.MT[0] = seed
	// 对数组的其它元素进行初始化
	for i := 1; i < 624; i++ {
		t := 1812433253*(mt.MT[i-1]^(mt.MT[i-1]>>30)) + i
		mt.MT[i] = t & 0xffffffff // 取最后的32位赋给MT[i]
	}
	return mt
}

func (mt *MersenneTwister) generate() {
	for i := 0; i < 624; i++ {
		// 2^31 = 0x80000000
		// 2^31-1 = 0x7fffffff
		y := (mt.MT[i] & 0x80000000) + (mt.MT[(i+1)%624] & 0x7fffffff)
		mt.MT[i] = mt.MT[(i+397)%624] ^ (y >> 1)
		if y&1 == 1 {
			mt.MT[i] ^= 2567483615
		}
	}
}

func (mt *MersenneTwister) Rand() int {
	if mt.index == 0 {
		mt.generate()
	}
	y := mt.MT[mt.index]
	y = y ^ (y >> 11)                // y右移11个bit
	y = y ^ ((y << 7) & 2636928640)  // y左移7个bit与2636928640相与，再与y进行异或
	y = y ^ ((y << 15) & 4022730752) // y左移15个bit与4022730752相与，再与y进行异或
	y = y ^ (y >> 18)                // y右移18个bit再与y进行异或
	mt.index = (mt.index + 1) % 624
	return y
}
