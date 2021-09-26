package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)

	i := 1
	o := false
	for c*i <= b {
		if c*i >= a {
			fmt.Printf("%d\n", c*i)
			o = true
			break
		}
		i++
	}
	if !o {
		fmt.Printf("-1\n")
	}
}
