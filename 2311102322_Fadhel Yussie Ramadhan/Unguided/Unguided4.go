package main

import "fmt"

func main() {
	var N int
	fmt.Print("masukkan bilangan bulat positif: ")
	fmt.Scanln(&N)

	fmt.Println("keluaran:")
	printSequence(N)
	fmt.Println()
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