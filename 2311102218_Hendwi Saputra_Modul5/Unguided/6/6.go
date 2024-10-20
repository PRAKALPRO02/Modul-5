package main

import "fmt"

func Perpangkatan(bil_dasar, pangkat int) int {
	if pangkat == 0 {
		return 1
	} else {
		return bil_dasar * Perpangkatan(bil_dasar, pangkat-1)
	}
}

func main() {
	var bil_dasar, pangkat int

	fmt.Print("Masukkan bilangan dan pangkatnya: ")
	fmt.Scan(&bil_dasar, &pangkat)

	hasil := Perpangkatan(bil_dasar, pangkat)
	fmt.Print(hasil)

}
