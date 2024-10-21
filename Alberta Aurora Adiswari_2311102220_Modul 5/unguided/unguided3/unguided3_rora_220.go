package main

import (
	"fmt"
)

func cetakFaktor(N, i int) {
	if i > N {
		return
	}
	if N%i == 0 {
		fmt.Print(i, " ")
	}
	cetakFaktor(N, i+1)
}

func main() {
	var N int

	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&N)

	cetakFaktor(N, 1)
	fmt.Println()
}
