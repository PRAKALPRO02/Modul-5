package main

import "fmt"

func isFaktor(bilangan, faktor int) bool {
	return bilangan%faktor == 0
}

func cetakFaktor(bilangan, faktor int) {
	if faktor > bilangan {
		return
	}
	if isFaktor(bilangan, faktor) {
		fmt.Print(faktor, " ")
	}
	cetakFaktor(bilangan, faktor+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	fmt.Printf("Faktor dari %d adalah: ", n)
	cetakFaktor(n, 1)
	fmt.Println()
}

// Zahrina Antika Malahati_2311102109
