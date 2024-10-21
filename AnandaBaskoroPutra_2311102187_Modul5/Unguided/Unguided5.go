package main

import "fmt"

func cetakBilanganGanjil(n int) {
	if n < 1 {
		return
	}
	cetakBilanganGanjil(n - 2)
	if n%2 != 0 {
		fmt.Print(n, " ")
	}
}

func main() {
	var N int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scanln(&N)

	fmt.Print("Keluaran : ")
	cetakBilanganGanjil(N)
	fmt.Println()
}