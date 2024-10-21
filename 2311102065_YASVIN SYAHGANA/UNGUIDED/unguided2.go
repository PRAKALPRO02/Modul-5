package main

import "fmt"

func polbint(n int) {
	if n == 0 {
		return
	}
	polbint(n - 1)
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}

	fmt.Println()
}

func main() {
	var n int
	fmt.Print("Input (n): ")
	fmt.Scan(&n)
	polbint(n)
}
