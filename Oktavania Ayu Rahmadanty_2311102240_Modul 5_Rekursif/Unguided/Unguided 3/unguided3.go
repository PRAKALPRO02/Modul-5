package main

import "fmt"

// Fungsi rekursif untuk mencari faktor-faktor bilangan
func faktorRekursif(n, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	faktorRekursif(n, i+1)
}

func main() {
	var N int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&N)

	fmt.Print("Faktor dari ", N, ": ")
	faktorRekursif(N, 1)
	fmt.Println()
}
