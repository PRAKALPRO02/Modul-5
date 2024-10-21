package main

import "fmt"

func Perpangkatan(bilangan, pangkat int) int {
	if pangkat == 0 {
		return 1
	}
	return bilangan * Perpangkatan(bilangan, pangkat-1)
}

func main() {
	var bilangan, pangkatan int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bilangan)
	fmt.Print("Masukkan pangkat: ")
	fmt.Scan(&pangkatan)
	hasil := Perpangkatan(bilangan, pangkatan)
	fmt.Println("Hasil:", hasil)
}
