package main

import "fmt"

func main() {
	var o int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&o)

	fmt.Print("Bilangan ganjil: ")
	bilanganGanjil(o)
	fmt.Print()
}

func bilanganGanjil(n int) {
	if n <= 0 {
		return
	}
	bilanganGanjil(n - 1)

	if n%2 != 0 {
		fmt.Print(n, " ")
	}
}
