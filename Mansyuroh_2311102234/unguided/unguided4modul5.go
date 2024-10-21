package main

import "fmt"

func cetakTurun234(angka int) {
	if angka < 1 {
		return
	}
	fmt.Print(angka, " ")
	cetakTurun234(angka - 1)
}

func cetakNaik234(batas, sekarang int) {
	if sekarang > batas {
		return
	}
	fmt.Print(sekarang, " ")
	cetakNaik234(batas, sekarang+1)
	fmt.Println()
}

func main() {
	var inputAngka234 int
	fmt.Print("\nMasukkan angka : ")
	fmt.Scan(&inputAngka234)
	fmt.Println()
	fmt.Println("Barisan bilangan : \n")
	cetakTurun234(inputAngka234)
	cetakNaik234(inputAngka234, 2)
	fmt.Println()
}
