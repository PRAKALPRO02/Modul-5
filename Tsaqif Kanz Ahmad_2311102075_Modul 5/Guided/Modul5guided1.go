package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n) // membaca input pengguna
	baris(n)     //memanggil fungsi rekursif
}

func baris(bilangan int) {
	if bilangan == 1 { //base case : jika bilangan sama dengan
		fmt.Println(1) // cetak angka 1
	} else { // jika bilangan lebih besar dari
		fmt.Println(bilangan) // cetak bilangan saat ini
		baris(bilangan - 1)   // panggil fungsi baris dengan
	}
}
