package main

import "fmt"

func printPattern(N int) {
	if N == 0 {
		return
	}
	fmt.Printf("%d ", N)
	if N > 1 {
		printPattern(N - 1)
		fmt.Printf("%d ", N)
	}
}

func main() {
	var N int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)

	printPattern(N)
}
