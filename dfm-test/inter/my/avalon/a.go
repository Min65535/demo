package avalon

import (
	"fmt"
)

func searchInts(a []int, x int) int {
	// This function body is a manually inlined version of:
	//
	//   return sort.Search(len(a), func(i int) bool { return a[i] > x }) - 1
	//
	// With better compiler optimizations, this may not be needed in the
	// future, but at the moment this change improves the go/printer
	// benchmark performance by ~30%. This has a direct impact on the
	// speed of gofmt and thus seems worthwhile (2011-04-29).
	// TODO(gri): Remove this when compilers have caught up.
	i, j := 0, len(a)
	for i < j {
		h := i + (j-i)/2 // avoid overflow when computing h
		fmt.Println("h:", h)
		fmt.Println("a[h]:", a[h])

		// i ≤ h < j
		if a[h] <= x {
			i = h + 1
		} else {
			j = h
		}
		fmt.Println("j:", j)
		fmt.Println("i:", i)
		fmt.Println("--------")
	}
	return i - 1
}

func binarySearch(a []int, x int) int {

	i, j := 0, len(a)

	for {
		h := i + (j-i)/2
		fmt.Println("h:", h)
		fmt.Println("a[h]:", a[h])
		if a[h] < x {
			i = h + 1
		} else if a[h] > x {
			j = h
		} else {
			return h
		}

		fmt.Println("j:", j)
		fmt.Println("i:", i)
		fmt.Println("--------")
	}
	// return i - 1
}
