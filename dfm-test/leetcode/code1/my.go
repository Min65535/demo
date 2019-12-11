package code1

import (
	"sort"
	"fmt"
)

const (
	A = 1
	T = 10
	J = 11
	Q = 12
	K = 13
	X = 0
)

const (
	Spades   = "s"
	Hearts   = "h"
	Diamonds = "d"
	Clubs    = "c"
	Ghost    = "n"
)

const (
	HighCard      = 1 << iota //单张大牌
	OnePair                   //一对
	TwoPairs                  //两对
	ThreeOfAKind              //三条
	Straight                  //顺子
	Flush                     //同花(5张牌)
	FullHouse                 //满堂彩(葫芦,三带二)
	FourOfAKind               //四条
	StraightFlush             //同花顺
	RoyalFlush                //皇家同花顺
)

//卡牌类型
type Group struct {
	Face  uint
	Color string
}

func (g Group) String() string {
	return fmt.Sprintf("Face: %d, Color: %s", g.Face, g.Color)
}

type Groups []Group

func (g Groups) Len() int {
	return len(g)
}

func (g Groups) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func (g Groups) Less(i, j int) bool {
	return g[i].Face < g[j].Face
}

//func Sort(arr []Group)  {
//	sort.Sort(Groups(arr))
//}

//给牌组排序
func SortTheCardsGroupArr(param []Group) (data []Group) {
	var i, j int
	var temp Group
	for i = 1; i < len(param); i++ {
		for j = i; j > 0; j-- {
			if param[j].Face < param[j-1].Face {
				temp = param[j]
				param[j] = param[j-1]
				param[j-1] = temp
			} else {
				break
			}
		}

	}
	data = param
	return
}


func GetInt(arr [] int) int {
	sort.Ints(arr)
	fmt.Println("arr:", arr)
	return arr[len(arr)/2]
}
