package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {
	var N string
	fmt.Scanf("%s", &N)

	rs := []rune(N)
	sort.Slice(rs, func(i, j int) bool {
		return rs[i] > rs[j]
	})

	var ans int64
	for bits := 0; bits < (1 << len(N)); bits++ {
		var A, B int64
		for j := 0; j < len(N); j++ {
			if bits&(1<<j) > 0 {
				a, _ := strconv.ParseInt(string(rs[j]), 10, 64)
				A = A*10 + a
			} else {
				b, _ := strconv.ParseInt(string(rs[j]), 10, 64)
				B = B*10 + b
			}
		}
		ans = int64(math.Max(float64(ans), float64(A*B)))
	}
	fmt.Printf("%v\n", ans)
}
