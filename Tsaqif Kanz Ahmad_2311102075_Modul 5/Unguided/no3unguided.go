package main

import (
	"fmt"
)

func findFactorsLoop(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Print(i, " ")
			count++
		}
	}
	return count
}

func main() {
	var N int

	fmt.Print("Masukkan angka N : ")
	fmt.Scan(&N)

	fmt.Print("Faktor dari ", N, " : ")
	totalFactors := findFactorsLoop(N)
	fmt.Println()
	fmt.Println("Total faktor = ", totalFactors)
}
