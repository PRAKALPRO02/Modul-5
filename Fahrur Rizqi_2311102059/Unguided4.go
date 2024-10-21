package main

import "fmt"

func printSequence(n int) {
	if n == 1 {
		fmt.Print(n, " ")
		return
	}
	fmt.Print(n, " ")
	//fahrur059
	printSequence(n - 1)
	fmt.Print(n, " ")
}

func main() {
	var N int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scanln(&N)

	fmt.Println("Keluaran:")
	printSequence(N)
	fmt.Println()
}
