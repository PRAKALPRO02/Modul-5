package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n) // Membaca input pengguna
	baris(n) // memanggil fungsi rekursif 'baris'
}

func baris(bilangan int) {
	if bilangan == 1 {
		fmt.Println(1)
	} else {
		fmt.Println(bilangan)
		baris(bilangan - 1)
	}
}