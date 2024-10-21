package main

import "fmt"

func cetakUrutan(jumlah int) {
	if jumlah == 1 {
		fmt.Print(jumlah, " ")
		return
	}
	fmt.Print(jumlah, " ")
	cetakUrutan(jumlah - 1)
	fmt.Print(jumlah, " ")
}

func main() {
	var jumlah int
	fmt.Print("Masukkan jumlah bilangan: ")
	fmt.Scanln(&jumlah)

	fmt.Print("Keluaran : ")
	cetakUrutan(jumlah)
	fmt.Println()
}