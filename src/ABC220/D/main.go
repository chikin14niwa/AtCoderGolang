package main

import (
	"fmt"
	"math"
	"strconv"
)

const k_div = 998244353

var nap map[string][]int

func main() {
	nap = map[string][]int{}
	var n int
	fmt.Scanf("%d", &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	rf := F(a, int(math.Floor(float64(n)/2)))
	rg := G(a, int(math.Floor(float64(n)/2)))
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", (rf[i]+rg[i])%k_div)
	}

}

func F(a []int, m int) []int {
	if m == len(a) {
		key := ""
		for i := 0; i < m; i++ {
			key += strconv.Itoa(a[i])
		}
		for k, v := range nap {
			fmt.Printf("F(%s) k(%s) => v(%v)\n", key, k, v)
		}
		if v, ok := nap[key]; ok {
			fmt.Printf("key %s already registered.", key)
			return v
		}
	}
	t := make([]int, len(a)-1)
	for i := 0; i < len(t); i++ {
		if i == 0 {
			t[i] = (a[0] + a[1]) % 10
		} else {
			t[i] = a[i+1]
		}
	}
	r := make([]int, 10)
	if len(t) == 1 {
		r[t[0]] = 1
		return r
	} else {
		rf := F(t, m)
		rg := G(t, m)
		for i := 0; i < 10; i++ {
			r[i] = (rf[i] + rg[i]) % k_div
		}
		if len(t) == m {
			key := ""
			for i := 0; i < m; i++ {
				key += strconv.Itoa(t[i])
			}
			fmt.Printf("F key:%s\n", key)
			nap[key] = r
		}
		return r
	}
}

func G(a []int, m int) []int {
	if m == len(a) {
		key := ""
		for i := 0; i < m; i++ {
			key += strconv.Itoa(a[i])
		}
		for k, v := range nap {
			fmt.Printf("G(%s) k(%s) => v(%v)\n", key, k, v)
		}
		if v, ok := nap[key]; ok {
			fmt.Printf("key %s already registered.", key)
			return v
		}
	}
	t := make([]int, len(a)-1)
	for i := 0; i < len(t); i++ {
		if i == 0 {
			t[i] = (a[0] * a[1]) % 10
		} else {
			t[i] = a[i+1]
		}
	}
	r := make([]int, 10)
	if len(t) == 1 {
		r[t[0]] = 1
		return r
	} else {
		rf := F(t, m)
		rg := G(t, m)
		for i := 0; i < 10; i++ {
			r[i] = (rf[i] + rg[i]) % k_div
		}
		if len(t) == m {
			key := ""
			for i := 0; i < m; i++ {
				key += strconv.Itoa(t[i])
			}
			fmt.Printf("G key:%s\n", key)
			nap[key] = r
		}
		return r
	}
}
