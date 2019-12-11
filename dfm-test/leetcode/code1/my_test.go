package code1

import (
	"testing"
	"fmt"
	"sort"
)

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
}
