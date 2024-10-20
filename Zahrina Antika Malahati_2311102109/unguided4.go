package main

import "fmt"

// Fungsi rekursif untuk menampilkan angka dari N hingga 1
func printDescending(n int) {
	if n > 0 {
		fmt.Printf("%d ", n)
		printDescending(n - 1)
	}
}

// Fungsi rekursif untuk menampilkan angka dari 1 hingga N
func printAscending(n, current int) {
	if current <= n {
		fmt.Printf("%d ", current)
		printAscending(n, current+1)
	}
}

// Fungsi utama yang menggabungkan kedua pola
func printPattern(n int) {
	printDescending(n)
	printAscending(n, 2)
	fmt.Println()
}

func main() {
	var n int

	// Contoh input
	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&n)

	// Cetak pola bilangan sesuai dengan input
	printPattern(n)
}

// Zahrina Antika Malahati_2311102109
