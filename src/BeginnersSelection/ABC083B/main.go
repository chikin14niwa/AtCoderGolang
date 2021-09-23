package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, a, b int
	fmt.Scanf("%d %d %d", &n, &a, &b)

	var s int64
	s = 0
	for i := 1; i <= n; i++ {
		nstr := strconv.Itoa(i)
		sum_digit := 0
		for j := 0; j < len(nstr); j++ {
			d, _ := strconv.ParseInt(string(nstr[j]), 10, 64)
			sum_digit = sum_digit + int(d)
		}
		if a <= sum_digit && sum_digit <= b {
			s += int64(i)
		}
	}

	fmt.Printf("%d\n", s)
}
