package main

import (
	"fmt"
)

// Fungsi rekursif untuk menampilkan bilangan ganjil
func tampilkanGanjil(n int) {
	if n <= 0 {
		return
	}
	tampilkanGanjil(n - 2)
	fmt.Printf("%d ", n)
}

func main() {
	var N int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)

	// Menyesuaikan agar N hanya bilangan ganjil
	if N%2 == 0 {
		N = N - 1
	}

	fmt.Printf("Bilangan ganjil dari 1 hingga %d: ", N)
	tampilkanGanjil(N)
	fmt.Println()
}

// Zahrina Antika Malahati_2311102109
