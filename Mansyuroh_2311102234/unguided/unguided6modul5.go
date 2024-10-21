package main

import "fmt"

func hitungPangkat234(basis, eksponen int) int {
	if eksponen == 0 {
		return 1
	}
	return basis * hitungPangkat234(basis, eksponen-1)
}

func main() {
	var basis234, pangkat234 int
	fmt.Print("\nMasukkan basis : ")
	fmt.Scan(&basis234)
	fmt.Print("Masukkan pangkat : ")
	fmt.Scan(&pangkat234)
	hasil := hitungPangkat234(basis234, pangkat234)
	fmt.Println("\nHasil pangkat : ", hasil)
	fmt.Println()
}
