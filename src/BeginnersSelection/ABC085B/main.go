package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	d := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &d[i])
	}

	sort.Ints(d)
	x := 0
	beforeD := 101
	for i := n - 1; i >= 0; i-- {
		if beforeD > d[i] {
			x++
			beforeD = d[i]
		}
	}
	fmt.Printf("%d\n", x)
}
