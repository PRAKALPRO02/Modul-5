package main

import "fmt"

func printOdd(N int) {
	if N <= 0 {
		return
	}
	if N%2 == 0 {
		printOdd(N - 1)
	} else {
		printOdd(N - 2)
		fmt.Printf("%d ", N)
	}
}

func main() {
	var N int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)

	printOdd(N)
}
