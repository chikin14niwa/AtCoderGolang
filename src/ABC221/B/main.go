package main

import (
	"fmt"
	"os"
)

func main() {
	var S, T string
	fmt.Scanf("%s\n%s", &S, &T)

	if S == T {
		fmt.Printf("Yes\n")
		os.Exit(0)
	}
	for i := 0; i < len(S)-1; i++ {
		s := switchString(S, i)
		if s == T {
			fmt.Printf("Yes\n")
			os.Exit(0)
		}
	}
	fmt.Printf("No\n")
}

func switchString(s string, i int) string {
	rs := []rune(s)
	str := ""
	if i > 0 {
		str += string(rs[:i])
	}
	str += string(rs[i+1]) + string(rs[i])
	if i+1 < len(s) {
		str += string(rs[i+2:])
	}
	return str
}
