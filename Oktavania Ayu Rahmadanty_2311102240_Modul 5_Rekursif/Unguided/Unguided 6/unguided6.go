package main

import (
	"fmt"
)

// Fungsi rekursif untuk menghitung x pangkat y (x^y)
func power(x, y int) int {
	if y == 0 {
		return 1 // Basis kasus: x^0 = 1
	}
	return x * power(x, y-1) // Rekursif: x * (x^(y-1))
}

func main() {
	// Input contoh dari soal
	var x1, y1 = 2, 2 // Contoh input 1
	var x2, y2 = 5, 3 // Contoh input 2

	// Cetak hasil perpangkatan
	fmt.Printf("%d^%d = %d\n", x1, y1, power(x1, y1)) // Output: 4
	fmt.Printf("%d^%d = %d\n", x2, y2, power(x2, y2)) // Output: 125
}
