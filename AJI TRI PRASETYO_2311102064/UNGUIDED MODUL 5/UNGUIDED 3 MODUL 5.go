package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Printf("Faktor dari %d adalah:\n", n)
	cetakFaktor(n)
}

func cetakFaktor(n int) {
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Print(i, " ")
		}
	}
}
