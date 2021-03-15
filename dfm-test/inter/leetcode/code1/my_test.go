package code1

import (
	"fmt"
	"sort"
	"testing"
)

type ArrSt struct {
	Index int `json:"index"`
	Num   int `json:"num"`
}

func min(input map[int]ArrSt) (data int) {

	return
}

func mySort(arr []int) (data []int) {
	my := make(map[int]ArrSt)
	for i := range arr {
		if _, ok := my[i]; ok {
			//if my[i].Num>
		}
		my[i] = ArrSt{i, arr[i]}
	}
	return
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var arr []int
	arr = append(arr, nums1...)
	arr = append(arr, nums2...)
	sort.Ints(arr)
	length := len(arr)
	if length%2 == 1 {
		return float64(arr[(length+1)/2-1])
	}
	return float64(arr[length/2-1]+arr[length/2]) / 2
}

func TestGetInt(t *testing.T) {
	arr := []int{1, 2, 3, 2, 2, 2, 1}
	fmt.Println(GetInt(arr))

	cards := []Group{
		{HighCard, Spades},
		{FullHouse, Hearts},
		{RoyalFlush, Clubs},
		{ThreeOfAKind, Diamonds},
	}

	fmt.Println(cards)

	sort.Sort(Groups(cards))
	fmt.Println(cards)

	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Face > cards[j].Face
	})
	fmt.Println(cards)

	arr1 := []int{1, 2}
	arr2 := []int{3, 4}
	num := findMedianSortedArrays(arr1, arr2)

	fmt.Println("num:", num)

}
