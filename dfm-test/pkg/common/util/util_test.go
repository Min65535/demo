package util

import (
	"fmt"
	"testing"
)

func TestCatch(t *testing.T) {
	fmt.Println("sss")
	defer Catch("www")
	panic("11111")

}
