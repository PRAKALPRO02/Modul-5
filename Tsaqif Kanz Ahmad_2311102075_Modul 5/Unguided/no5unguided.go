package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scanln(&n)

	fmt.Println("Barisan bilangan ganjil:")
	printOddNumbers(n)
}

func printOddNumbers(n int) {
	if n >= 1 {
		printOddNumbers(n - 1)
		if n%2 != 0 {
			fmt.Print(n, " ")
		}
	}
}
