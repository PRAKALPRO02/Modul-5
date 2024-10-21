package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n) // Membaca input pengguna
	baris(n)     // Memanggil fungsi rekursif 'baris'
}

func baris(bilangan int) {
	if bilangan == 1 { // Base case: Jika bilangan sama dengan 1
		fmt.Println(1) // Cetak angka 1
	} else { // Jika bilangan lebih besar dari 1
		fmt.Println(bilangan) // cetak bilangan saat ini
		baris(bilangan - 1)   // Panggil fungsi 'baris' dengan bilangan
	}
}
