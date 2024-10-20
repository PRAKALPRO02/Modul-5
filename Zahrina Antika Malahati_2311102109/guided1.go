package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n) // Membaca input pengguna
	baris(n)     // Memanggil fugsi rekusif baris
}

func baris(bilangan int) {
	if bilangan == 1 { // Base case: Jika bilangan sama dengan 1
		fmt.Println(1) // Cetak angka 1
	} else {
		fmt.Println(bilangan)
		baris(bilangan - 1)
	}
}
