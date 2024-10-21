package main

import (
	"fmt"
)

func faktor_219(n, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	faktor_219(n, i+1)
}

func main() {
	var n int
	fmt.Scan(&n)

	faktor_219(n, 1)
	fmt.Println()
}
