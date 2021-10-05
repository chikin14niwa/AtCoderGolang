package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B int
	fmt.Scanf("%d %d", &A, &B)
	fmt.Printf("%d\n", int64(math.Pow(float64(32), float64(A-B))))
}
