package main

import "fmt"

func printStars(n, baris int) {
	if n == 0 {
		return
	}
	for i := 0; i < baris; i++ {
		fmt.Print("*")
	}
	fmt.Println()
	printStars(n-1, baris+1)
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&n)

	printStars(n, 1)
}

// Zahrina Antika Malahati_2311102109
