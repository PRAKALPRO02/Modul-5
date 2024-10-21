package main

import "fmt"

func cetakFaktor(n, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	cetakFaktor(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)

	fmt.Println("Faktor dari", n, "adalah:")
	cetakFaktor(n, 1)
	fmt.Println()
	return
}
