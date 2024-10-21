package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n) // memaca input pengguna
	baris(n)     // memanggil fungsi rekursif 'baris'
}

func baris(bilangan int) {
	if bilangan == 1 { // base case: Jika bilangan sama dengan 1
		fmt.Println(1) // cetak angka 1
	} else { // jika bilangan lebih besar dari 1
		fmt.Println(bilangan) // cetak bilangan saat ini
		baris(bilangan - 1)
	}
}
