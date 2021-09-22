package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d\n", &a, &b)

	var s string
	if a*b%2 == 0 {
		s = "Even"
	} else {
		s = "Odd"
	}
	fmt.Printf("%s\n", s)
}
