package main

import "fmt"

func tampilkanGanjil234(angka, batas int) {
	if angka > batas {
		return
	}
	if angka%2 != 0 {
		fmt.Print(angka, " ")
	}
	tampilkanGanjil234(angka+1, batas)
}

func main() {
	var maxBilangan234 int
	fmt.Print("\nMasukkan jumlah bilangan: ")
	fmt.Scan(&maxBilangan234)
	fmt.Println("\nBilangan yang ganjil : ")
	tampilkanGanjil234(0, maxBilangan234)
	fmt.Println()
}
