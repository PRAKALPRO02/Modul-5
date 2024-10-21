package main

import (
	"fmt"
)

func cetakBintang(n int) {
	if n == 0 {
		return
	}
	fmt.Print("*")
	cetakBintang(n - 1)
}

func cetakPola(n int) {
	if n == 0 {
		return
	}
	cetakPola(n - 1)
	cetakBintang(n)
	fmt.Println()
}

func main() {
	var N int

	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&N)

	cetakPola(N)
}
