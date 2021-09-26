package main

import (
	"fmt"
	"strconv"
)

func main() {
	var k int
	var a, b string

	fmt.Scanf("%d\n%s %s", &k, &a, &b)
	ai, _ := strconv.ParseInt(a, k, 64)
	bi, _ := strconv.ParseInt(b, k, 64)
	fmt.Printf("%d\n", ai*bi)
}
