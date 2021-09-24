package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var N, Y int
	fmt.Scanf("%d %d", &N, &Y)

	var x, y, z int
	maxXn := int(math.Floor(float64(Y) / 10000))
	for x = maxXn; x >= 0; x-- {
		if 10000*x > Y {
			continue
		}
		for y = N - x; y >= 0; y-- {
			if 10000*x+5000*y > Y {
				continue
			}
			for z = N - x - y; z >= 0; z-- {
				if 10000*x+5000*y+1000*z > Y {
					continue
				}
				if x+y+z == N && 10000*x+5000*y+1000*z == Y {
					fmt.Printf("%d %d %d\n", x, y, z)
					os.Exit(0)
				}
			}
		}
	}
	fmt.Printf("-1 -1 -1\n")
}
