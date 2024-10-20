package main

import (
	"fmt"
)

func printSequence(n int) {
	for current := 1; current <= n; current++ {
		fmt.Print(current, " ")
	}
	for current := n - 1; current >= 1; current-- {
		fmt.Print(current, " ")
	}
}

func main() {
	var N int
	
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&N)

	printSequence(N)
	fmt.Println() 
}
