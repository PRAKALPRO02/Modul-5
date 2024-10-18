package main

import "fmt"

func deretFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return deretFibonacci(n-1) + deretFibonacci(n-2)
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(deretFibonacci(n))
}