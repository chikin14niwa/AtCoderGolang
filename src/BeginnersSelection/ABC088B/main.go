package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	sort.Ints(a)

	alice := 0
	bob := 0
	aliceTurn := (n - 1) % 2
	for i := n - 1; i >= 0; i-- {
		if i%2 == aliceTurn {
			alice += a[i]
		} else {
			bob += a[i]
		}
	}
	fmt.Printf("%d\n", alice-bob)
}
