package main

import (
	"fmt"
	"math"
	"math/big"
	"sort"
	"strconv"
)

func main() {
	var N string
	fmt.Scanf("%s", &N)

	tmp := []rune(N)
	a := make([]int, len(N))
	for i := 0; i < len(N); i++ {
		a[i], _ = strconv.Atoi(string(tmp[i]))
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	str := ""
	for i := 0; i < len(N); i++ {
		str += strconv.Itoa(a[len(N)-i-1])
	}
	sortN := []rune(str)
	n := int64(math.Floor(float64(len(N) / 2)))
	var ans int64
	A := maxNumFromString(string(sortN[len(sortN)-1]))
	B := maxNumFromString(string(sortN[:len(sortN)-1]))
	ans = A * B
	if n > 1 {
		A = maxNumFromString(string(sortN[len(sortN)-1]) + string(sortN[len(sortN)-3]))
		B = maxNumFromString(string(sortN[:len(sortN)-3]) + string(sortN[len(sortN)-2]))
		bigA := big.NewFloat(float64(A))
		bigB := big.NewFloat(float64(B))
		bigAns, _ := bigA.Mul(bigA, bigB).Float64()
		ans = int64(math.Max(float64(ans), bigAns))
	}
	if n > 2 {
		A = maxNumFromString(string(sortN[len(sortN)-1]) + string(sortN[len(sortN)-3]) + string(sortN[len(sortN)-5]))
		B = maxNumFromString(string(sortN[:len(sortN)-5]) + string(sortN[len(sortN)-4]) + string(sortN[len(sortN)-2]))
		bigA := big.NewFloat(float64(A))
		bigB := big.NewFloat(float64(B))
		bigAns, _ := bigA.Mul(bigA, bigB).Float64()
		ans = int64(math.Max(float64(ans), bigAns))
	}
	if n > 3 {
		A = maxNumFromString(string(sortN[len(sortN)-1]) + string(sortN[len(sortN)-3]) + string(sortN[len(sortN)-5]) + string(sortN[len(sortN)-7]))
		B = maxNumFromString(string(sortN[:len(sortN)-7]) + string(sortN[len(sortN)-6]) + string(sortN[len(sortN)-4]) + string(sortN[len(sortN)-2]))
		bigA := big.NewFloat(float64(A))
		bigB := big.NewFloat(float64(B))
		bigAns, _ := bigA.Mul(bigA, bigB).Float64()
		ans = int64(math.Max(float64(ans), bigAns))
	}
	fmt.Printf("%d\n", ans)
}

func maxNumFromString(s string) int64 {
	rs := []rune(s)
	a := make([]int64, len(s))
	for i := 0; i < len(s); i++ {
		a[i], _ = strconv.ParseInt(string(rs[i]), 10, 64)
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	var sum int64
	for i := 0; i < len(a); i++ {
		sum += a[i] * int64(math.Pow10(i))
	}
	return sum
}
