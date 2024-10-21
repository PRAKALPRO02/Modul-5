package main

import "fmt"

func Fibonacci_219(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci_219(n-1) + Fibonacci_219(n-2)
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(Fibonacci_219(n))
}
