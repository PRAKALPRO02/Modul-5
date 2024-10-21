package main

import "fmt"

func findFactors(n, divisor int) {
	if divisor > n {
		return
	}

	if n%divisor == 0 {
		fmt.Print(divisor, " ")
	}

	findFactors(n, divisor+1)
}

func main() {
	var n int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	fmt.Println("Faktor-faktor dari", n, "adalah:")
	findFactors(n, 1)
}
