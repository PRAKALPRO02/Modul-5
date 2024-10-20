package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)
	fibonacciSeries(n)
}

func fibonacciSeries(n int) {
	a, b := 0, 1
	for i := 0; i <= n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println()
}
