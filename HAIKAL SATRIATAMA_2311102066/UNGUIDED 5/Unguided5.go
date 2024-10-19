package main

import "fmt"

func tampilkanGanjil(n int) {
	for i := 1; i <= n; i += 2 {
		fmt.Print(i, " ")
	}
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&n)
	tampilkanGanjil(n)
}
