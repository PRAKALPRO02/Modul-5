package main

import "fmt"

// Fungsi untuk mencetak faktor n
func printFactors(n, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	printFactors(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan : ")
	fmt.Scan(&n)
	fmt.Print("Faktor dari ", n, " adalah: ")
	printFactors(n, 1)
	fmt.Println()
}
