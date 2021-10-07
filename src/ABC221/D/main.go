package main

import (
	"fmt"
	"sort"
)

type Point struct {
	P int64
	C int64
}

func main() {
	var N int64
	fmt.Scanf("%d", &N)
	A := make([]int64, N)
	B := make([]int64, N)
	points := make([]Point, 2*N)
	for i := 0; int64(i) < N; i++ {
		fmt.Scanf("%d %d", &A[i], &B[i])
		points[2*i] = Point{P: A[i], C: 1}
		points[2*i+1] = Point{P: A[i] + B[i], C: -1}
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i].P < points[j].P
	})
	ans := make([]int64, N+1)
	var cnt int64
	for i := 0; i < len(points)-1; i++ {
		cnt += points[i].C
		ans[cnt] += points[i+1].P - points[i].P
	}
	for i := 1; int64(i) <= N; i++ {
		fmt.Printf("%d ", ans[i])
	}
	fmt.Println()
}
