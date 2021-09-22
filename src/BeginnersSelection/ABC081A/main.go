package main

import "fmt"

func main() {
	var s string
	fmt.Scanf("%s", &s)

	i := 0
	for _, si := range s {
		if si == '1' {
			i += 1
		}
	}
	fmt.Printf("%d\n", i)
}
