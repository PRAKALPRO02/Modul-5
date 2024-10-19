package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(perjumlahan(n))
}

func perjumlahan(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return n + perjumlahan(n-1)
	}
}
