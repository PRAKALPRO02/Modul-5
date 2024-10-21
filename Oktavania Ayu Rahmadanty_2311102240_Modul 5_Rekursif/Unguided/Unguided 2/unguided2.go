package main

import "fmt"

// Fungsi rekursif untuk menampilkan satu baris bintang
func cetakBintang(n int) {
	if n == 0 {
		return
	}
	fmt.Print("*")
	cetakBintang(n - 1)
}

// Fungsi rekursif untuk menampilkan pola bintang bertingkat
func polaBintang(n, current int) {
	if current > n {
		return
	}
	cetakBintang(current)
	fmt.Println()
	polaBintang(n, current+1)
}

func main() {
	var N int
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&N)

	polaBintang(N, 1)
}
