package algorithm

// 梅森旋转算法 go
// Mersenne Twister

type MersenneTwister struct {
	index int
	MT    [MersenneTwisterN]int // 624 * 32 - 31 = 19937
}

const (
	MersenneTwisterInitOperand = 0x6c078965
	MersenneTwisterMaxBits     = 0xffffffff
	MersenneTwisterUpperBits   = 0x80000000
	MersenneTwisterLowerBits   = 0x7fffffff

	MersenneTwisterM = 397
	MersenneTwisterN = 624
	MersenneTwisterA = 0x9908b0df
	MersenneTwisterB = 0x9d2c5680
	MersenneTwisterC = 0xefc60000

	MaxValue = 4294967295
)

func NewMersenneTwister(seed int) *MersenneTwister {
	mt := &MersenneTwister{}
	mt.index = 0
	mt.MT[0] = seed
	// 对数组的其它元素进行初始化
	for i := 1; i < MersenneTwisterN; i++ {
		t := MersenneTwisterInitOperand*(mt.MT[i-1]^(mt.MT[i-1]>>30)) + i
		mt.MT[i] = t & MersenneTwisterMaxBits // 取最后的32位赋给MT[i]
	}
	return mt
}

func (mt *MersenneTwister) generate() {
	for i := 0; i < MersenneTwisterN; i++ {
		// 2^31 = 0x80000000
		// 2^31-1 = 0x7fffffff
		y := (mt.MT[i] & MersenneTwisterUpperBits) + (mt.MT[(i+1)%MersenneTwisterN] & MersenneTwisterLowerBits)
		mt.MT[i] = mt.MT[(i+MersenneTwisterM)%MersenneTwisterN] ^ (y >> 1)
		if y&1 == 1 {
			mt.MT[i] ^= MersenneTwisterA
		}
	}
}

// RandInt 从2^32-1取出随机整数
func (mt *MersenneTwister) RandInt() int {
	if mt.index == 0 {
		mt.generate()
	}
	y := mt.MT[mt.index]
	y = y ^ (y >> 11)                      // y右移11个bit
	y = y ^ ((y << 7) & MersenneTwisterB)  // y左移7个bit与2636928640相与,再与y进行异或
	y = y ^ ((y << 15) & MersenneTwisterC) // y左移15个bit与4022730752相与,再与y进行异或
	y = y ^ (y >> 18)                      // y右移18个bit再与y进行异或
	mt.index = (mt.index + 1) % MersenneTwisterN
	return y
}

// RandomFloat64Value 从[0,1)取出64位浮点数
func (mt *MersenneTwister) RandomFloat64Value() float64 {
	return float64(mt.RandInt()) / float64(MaxValue)
}

// RangeInt 从[min,max)取出随机整数
func (mt *MersenneTwister) RangeInt(min, max int) int {
	if max <= min {
		return min
	}
	return min + int(float64(max-min)*mt.RandomFloat64Value())
}
