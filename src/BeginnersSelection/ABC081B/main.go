package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	a := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	ans := 0
	for {
		isDead := false
		for i, v := range a {
			if v%2 == 0 {
				a[i] = v / 2
			} else {
				isDead = true
			}
		}
		if isDead {
			break
		}
		ans++
	}
	fmt.Printf("%d\n", ans)
}
