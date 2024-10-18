package main

import "fmt"

func main() {
	var N int
	fmt.Print("Masukkan bilangan bulat positif N : ")
	fmt.Scanln(&N)
	printSequence(N)
}

func printSequence(n int) {
	if n == 1 {
		fmt.Print(n, " ")
		return
	}
	fmt.Print(n, " ")
	printSequence(n - 1)
	fmt.Print(n, " ")
}
