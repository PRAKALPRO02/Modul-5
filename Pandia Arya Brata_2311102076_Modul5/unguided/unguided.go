package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i <= n; i++ {
		fmt.Printf("%d ", fibonaci(i))
	}
	fmt.Println()
}

func fibonaci(n int) int {
	if n <= 1 {
		return n
	} else {
		return fibonaci(n-1) + fibonaci(n-2)
	}
}
