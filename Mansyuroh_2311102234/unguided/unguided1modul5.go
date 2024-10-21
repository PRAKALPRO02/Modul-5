package main

import "fmt"

func hitungFibonacci234(angka234 int) int {
	if angka234 <= 1 {
		return angka234
	}
	return hitungFibonacci234(angka234-1) + hitungFibonacci234(angka234-2)
}

func main() {
	var batasDeret234 int
	indeks234 := 0

	fmt.Println("\n*PROGRAM MENGHITUNG DERET FIBONACCI*")
	fmt.Print("\nBatas deret: ")
	fmt.Scan(&batasDeret234)
	fmt.Println()

	for indeks := 0; indeks <= batasDeret234; indeks++ {
		fmt.Printf("n%d : %d\n", indeks, hitungFibonacci234(indeks234))
		indeks234++
	}

	fmt.Println()
}
