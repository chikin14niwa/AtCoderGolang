package main

import "fmt"

func main() {
	var a, b, c, x int
	fmt.Scanf("%d\n%d\n%d\n%d", &a, &b, &c, &x)

	ans := 0
	for i := 0; i <= a; i++ {
		if 500*i > x {
			break
		}
		for j := 0; j <= b; j++ {
			if 500*i+100*j > x {
				break
			}
			for k := 0; k <= c; k++ {
				if 500*i+100*j+50*k > x {
					break
				}
				if 500*i+100*j+50*k == x {
					ans++
					break
				}
			}
		}
	}

	fmt.Printf("%d\n", ans)
}
