package main

import (
	"fmt"
)

func printOddNumbers(N, i int) {

	if i > N {
		return
	}

	if i%2 != 0 {
		fmt.Print(i, " ")
	}

	printOddNumbers(N, i+1)
}

func main() {
	var N int

	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)
	fmt.Println("Bilangan ganjil dari 1 hingga", N, "adalah:")
	printOddNumbers(N, 1)
}
