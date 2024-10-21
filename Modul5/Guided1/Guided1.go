package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n) //Membaca input pengguna
	baris(n)     //Memanggil fungsi rekursif "baris"
}

func baris(bilangan int) {
	if bilangan == 1 { //Base case: jika bilangan sama dengan
		fmt.Println(1) //Cetak angka
	} else { //Jika bilangan lebih besar dari
		fmt.Println(bilangan) //Cetak bilanga saat ini
		baris(bilangan - 1)   //Panggil fungsi "baris" dengan
	}
}
