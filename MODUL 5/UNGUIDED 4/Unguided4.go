package main

import "fmt"

func printSequence(n, current int) {
	if current > n {
		return
	}
	fmt.Print(current, " ")
	printSequence(n, current+1)
	if current < n {
		fmt.Print(current, " ")
	}
}

func main() {
	var N int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&N)
	printSequence(N, 1)
	fmt.Println()
}
