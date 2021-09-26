package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
)

func main() {
	var n, asum int64
	fmt.Scanf("%d", &n)
	a := make([]int64, n)
	asum = 0
	var i, j int64
	for i = 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
		asum += a[i]
	}
	var s, x int64
	fmt.Scanf("%d", &x)
	s = 0
	maxi := big.NewInt(10)
	for i = 0; i < 100; i++ {
		maxi = maxi.Mul(maxi, big.NewInt(10))
	}
	loopi := big.NewInt(int64(math.Floor(float64(x)/float64(asum))) - 1)
	s = asum * loopi.Int64()
	for ; maxi.Cmp(loopi) == 1; loopi.Add(loopi, big.NewInt(1)) {
		for j = 0; j < n; j++ {
			s += a[j]
			if s > x {
				fmt.Printf("%s\n", loopi.Add(loopi.Mul(loopi, big.NewInt(n)), big.NewInt(j+1)).String())
				os.Exit(0)
			}
		}
	}
}
