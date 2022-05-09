package mersennetwister

import "github.com/dipperin/go-ms-toolkit/json"

// 梅森旋转算法 go
// Mersenne Twister

const (
	ALMersenneTwisterInitOperand = 0x6c078965
	ALMersenneTwisterMaxBits     = 0xffffffff
	ALMersenneTwisterUpperBits   = 0x80000000
	ALMersenneTwisterLowerBits   = 0x7fffffff

	ALMersenneTwisterM = 397
	ALMersenneTwisterN = 624
	ALMersenneTwisterA = 0x9908b0df
	ALMersenneTwisterB = 0x9d2c5680
	ALMersenneTwisterC = 0xefc60000

	ALMersenneTwisterMaxValue = 4294967295 // 2^32-1
)

type MersenneTwister struct {
	index int
	MT    [ALMersenneTwisterN]int // 624 * 32 - 31 = 19937
}

type Array []interface{}

func (a *Array) Unmarshal(res interface{}) error {
	return json.ParseJson(json.StringifyJson(a), res)
}

func Marshal(input interface{}) (Array, error) {
	var a Array
	return a, json.ParseJson(json.StringifyJson(input), &a)
}

func NewMersenneTwister(seed int) *MersenneTwister {
	mt := &MersenneTwister{}
	mt.index = 0
	mt.MT[0] = seed
	// 对数组的其它元素进行初始化
	for i := 1; i < ALMersenneTwisterN; i++ {
		// t := MersenneTwisterInitOperand*(mt.MT[i-1]^(mt.MT[i-1]>>30)) + i
		// mt.MT[i] = t & MersenneTwisterMaxBits // 取最后的32位赋给MT[i]
		pre := mt.MT[i-1]
		tem := ALMersenneTwisterInitOperand*(pre^(pre>>30)) + i
		mt.MT[i] = tem & ALMersenneTwisterMaxBits
	}
	return mt
}

func (mt *MersenneTwister) generate() {
	for i := 0; i < ALMersenneTwisterN; i++ {
		// 2^31 = 0x80000000
		// 2^31-1 = 0x7fffffff
		y := (mt.MT[i] & ALMersenneTwisterUpperBits) + (mt.MT[(i+1)%ALMersenneTwisterN] & ALMersenneTwisterLowerBits)
		mt.MT[i] = mt.MT[(i+ALMersenneTwisterM)%ALMersenneTwisterN] ^ (y >> 1)
		// if y&1 == 1 {
		// 	mt.MT[i] ^= MersenneTwisterA
		// }
		if y%2 > 0 {
			mt.MT[i] ^= ALMersenneTwisterA
		}
	}
}

// RandInt 从2^32-1取出随机整数
func (mt *MersenneTwister) RandInt() int {
	if mt.index == 0 {
		mt.generate()
	}
	y := mt.MT[mt.index]
	y = y ^ (y >> 11)                        // y右移11个bit
	y = y ^ ((y << 7) & ALMersenneTwisterB)  // y左移7个bit与2636928640相与,再与y进行异或
	y = y ^ ((y << 15) & ALMersenneTwisterC) // y左移15个bit与4022730752相与,再与y进行异或
	y = y ^ (y >> 18)                        // y右移18个bit再与y进行异或
	mt.index = (mt.index + 1) % ALMersenneTwisterN
	return y
}

// RandomFloat64Value 从[0,1)取出64位浮点数
func (mt *MersenneTwister) RandomFloat64Value() float64 {
	return float64(mt.RandInt()) / float64(ALMersenneTwisterMaxValue)
}

// RangeInt 从[min,max)取出随机整数
func (mt *MersenneTwister) RangeInt(min, max int) int {
	if max <= min {
		return min
	}
	return min + int(float64(max-min)*mt.RandomFloat64Value())
}

// RangeFloat64 从[min,max)取出随机浮点数
func (mt *MersenneTwister) RangeFloat64(min, max float64) float64 {
	if max <= min {
		return min
	}
	return min + (max-min)*mt.RandomFloat64Value()
}

// ShuffleArray 数组洗牌
func (mt *MersenneTwister) ShuffleArray(arr []interface{}) {
	length := len(arr)
	for i := 0; i < len(arr); i++ {
		ind := mt.RangeInt(0, length)
		tem := arr[ind]
		arr[ind] = arr[i]
		arr[i] = tem
	}
}

// GetRandomArray 从输入数组随机取出count个数元素数组
func (mt *MersenneTwister) GetRandomArray(arr []interface{}, count int) []interface{} {
	var res []interface{}
	if len(arr) > count {
		var indexes []int
		for i := 0; i < count; i++ {
			ind := mt.RangeInt(0, len(arr)-i)
			for j := 0; j < len(indexes); j++ {
				if ind >= indexes[j] {
					ind++
				}
			}
			indexes = append(indexes, ind)
		}
		for i := 0; i < len(indexes); i++ {
			ind2 := indexes[i]
			res = append(res, arr[ind2])
		}
	} else {
		res = append(res, arr...)
	}
	return res
}
