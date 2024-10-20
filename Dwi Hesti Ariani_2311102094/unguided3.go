package main

import (
	"fmt"
)

func printFactors(N, i int) {

	if i > N {
		return
	}

	if N%i == 0 {
		fmt.Println(i)
	}
	printFactors(N, i+1)
}
func main() {
	var N int

	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)

	fmt.Println("Faktor-faktor dari", N, "adalah:")
	printFactors(N, 1)
}
