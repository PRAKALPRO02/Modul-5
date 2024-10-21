package main

import "fmt"

func tampilkanGanjil(n, i int) {
	if i > n {
		return
	}
	if i%2 != 0 {
		fmt.Print(i, " ")
	}
	tampilkanGanjil(n, i+1)
}

func main() {
	var N int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&N)

	fmt.Print("Bilangan ganjil dari 1 hingga ", N, ": ")
	tampilkanGanjil(N, 1)
	fmt.Println()
}
