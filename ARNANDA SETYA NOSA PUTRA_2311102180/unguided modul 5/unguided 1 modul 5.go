package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)
	fibonacciSeries(n)
}

func fibonacciSeries(n int) {
	for i := 0; i <= n; i++ {
		fmt.Print(fibonacci(i), " ")
	}
	fmt.Println()
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
