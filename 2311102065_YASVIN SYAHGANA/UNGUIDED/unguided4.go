package main

import "fmt"

func domino(n, f int) {
	if n == 1 {
		fmt.Print(1, " ")
		return
	}

	fmt.Print(n, " ")
	domino(n-1, f)
	fmt.Print(n, " ")
}

func main() {
	var n int
	fmt.Print("Input (n): ")
	fmt.Scan(&n)
	domino(n, n)
}
