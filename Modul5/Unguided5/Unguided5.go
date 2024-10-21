package main

import "fmt"

func cetakBilangan(n int) {
	if n > 0 {
		cetakBilangan(n - 1)
		if n%2 != 0 {
			fmt.Printf("%d ", n)
		}
	}
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)
	fmt.Print("Keluaran barisan: ")
	cetakBilangan(n)
	fmt.Println()
}
