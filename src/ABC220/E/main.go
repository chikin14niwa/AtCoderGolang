package main

import "fmt"

func main() {
	var N, D, i, j, k, s int64
	fmt.Scanf("%d %d", &N, &D)

	dp := make([][]int64, N+1)
	for i = 0; i <= N; i++ {
		dp[i] = make([]int64, D+1)
	}
	dp[0][1] = 2

	for i = 1; i <= D; i++ {
		for j = 1; j <= N; j++ {
			if j+N > i {
				if N-j >= i {
					dp[j][i] = 2 ^ i
					fmt.Printf("dp[%d][%d] = %d\n", j, i, dp[j][i])
				}
				for k = 1; k <= j; k++ {
					fmt.Printf("i:%d, j:%d, k:%d\n", i, j, k)
					if i > k {
						dp[j][i] += dp[j-k][i-k] / 2
						fmt.Printf("dp[%d][%d] = %d\n", j, i, dp[j][i])
					}
				}
			}
		}
	}

	s = 0
	for i = 0; i <= N; i++ {
		s += dp[i][D]*2 ^ i
		s = s % 998244353
	}
	fmt.Printf("%d\n", s)
}
