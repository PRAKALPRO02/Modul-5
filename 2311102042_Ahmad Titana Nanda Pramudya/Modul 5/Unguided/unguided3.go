package main

import (
	"fmt"
)

func findFactors(n, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	findFactors(n, i+1)
}

func main() {
	var N int

	fmt.Print("Masukkan angka N: ")
	fmt.Scan(&N)

	
	fmt.Print("Faktor dari", N, ": ")
	findFactors(N, 1)
	fmt.Println() 
}
