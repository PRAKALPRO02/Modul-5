package main

import "fmt"

func tampilGanjil(n int, current int) {
	if current > n {
		return
	}
	fmt.Print(current, " ")
	tampilGanjil(n, current+2)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&n)
	tampilGanjil(n, 1)
}