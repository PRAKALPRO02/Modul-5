package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n) //membaca input pengguna
	baris(n)     //memanggil fungsi rekusrsif 'baris'
}

func baris(bilangan int) {
	if bilangan == 1 { //base case
		fmt.Println(1) //cetak angka 1
	} else {
		fmt.Println(bilangan)
		baris(bilangan - 1)
	}
}
