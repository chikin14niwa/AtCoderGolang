package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	A := make([]int, N+1)
	dp := make([][10]int64, N+1)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &A[i+1])
	}

	dp[1][A[1]] = 1
	for i := 1; i < N; i++ {
		for j := 0; j < 10; j++ {
			dp[i+1][(j+A[i+1])%10] = int64(math.Mod(float64(dp[i+1][(j+A[i+1])%10]+dp[i][j]), 998244353))
			dp[i+1][(j*A[i+1])%10] = int64(math.Mod(float64(dp[i+1][(j*A[i+1])%10]+dp[i][j]), 998244353))
		}
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", dp[N][i]%998244353)
	}
}
