package main

import (
	"fmt"
	"math"
)

func main() {
	var N, i, j, l int64
	fmt.Scanf("%d", &N)
	A := make([]int64, N)
	B := make([]int64, N)
	for i = 0; i < N; i++ {
		fmt.Scanf("%d %d", &A[i], &B[i])
		l = int64(math.Max(float64(l), float64(A[i]+B[i])))
	}
	dp := make([]int64, l+1)
	for i = 0; i < N; i++ {
		for j = 0; j < B[i]; j++ {
			dp[A[i]+j]++
		}
	}
	ans := make([]int64, N)
	for i = 1; i < l; i++ {
		ans[dp[i]-1]++
	}
	for i = 0; i < N; i++ {
		fmt.Printf("%d ", ans[i])
	}
	fmt.Println()
}
