package main

import (
	"fmt"
)

// Fungsi untuk menghitung pangkat
func hitungPangkat(x int, y int) int {
	if y < 0 {
		return 0 // Mengembalikan 0 untuk pangkat negatif, karena hasil pangkat bilangan bulat tidak terdefinisi
	}

	hasil := 1
	for i := 0; i < y; i++ {
		hasil *= x // Hitung pangkat dengan iterasi
	}
	return hasil
}

func main() {
	var bilangan int
	var pangkat int

	// Input dari pengguna
	fmt.Print("Masukkan bilangan: ")
	fmt.Scanln(&bilangan)
	fmt.Print("Masukkan pangkat: ")
	fmt.Scanln(&pangkat)

	// Menghitung hasil pangkat
	hasil := hitungPangkat(bilangan, pangkat)

	// Menampilkan hasil
	fmt.Printf("Hasil %d pangkat %d adalah %d\n", bilangan, pangkat, hasil)
}
