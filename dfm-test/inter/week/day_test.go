package week

import (
	"fmt"
	"testing"
	"time"
)

func TestGetSunday(t *testing.T) {
	GetSunday()

	fff := fmt.Sprintf("%d", time.Now().Unix())
	fmt.Println(fff, len(fff))
	//
	ss := time.Second
	fmt.Println(ss.Microseconds(), ss.Milliseconds())
}
