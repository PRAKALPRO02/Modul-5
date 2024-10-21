package main

import "fmt"

func tampilkanBintang234(hitung234 int) {
	if hitung234 <= 0 {
		return
	}
	fmt.Print("*")
	tampilkanBintang234(hitung234 - 1)
}

func tampilkanBaris234(totalBaris234, urutan234 int) {
	if urutan234 > totalBaris234 {
		return
	}
	tampilkanBintang234(urutan234)
	fmt.Println()
	tampilkanBaris234(totalBaris234, urutan234+1)
}

func main() {
	var jumlahBaris234 int
	fmt.Print("\nMasukkan jumlah baris(n) : ")
	fmt.Scan(&jumlahBaris234)
	fmt.Println("\nPola : \n")
	tampilkanBaris234(jumlahBaris234, 1)
	fmt.Println()
}
