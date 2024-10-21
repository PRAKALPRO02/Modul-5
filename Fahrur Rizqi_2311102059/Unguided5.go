package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scanln(&n)

	fmt.Println("Keluaran:")
	printOddNumbers(n)
}

func printOddNumbers(n int) {
	if n >= 1 {
		printOddNumbers(n - 2)
		//fahrur059
		fmt.Print(n, " ")
	}
}
