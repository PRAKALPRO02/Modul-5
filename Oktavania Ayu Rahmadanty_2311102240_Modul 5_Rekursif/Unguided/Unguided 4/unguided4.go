package main

import (
	"fmt"
)

func printSequence(n, current int) {
	fmt.Printf("%d ", current)

	if current > 1 {
		printSequence(n, current-1)
	}

	if current < n {
		fmt.Printf("%d ", current)
	}
}

func main() {
	fmt.Print("Masukan 5: ")
	printSequence(5, 5)
	fmt.Println()

	fmt.Print("Masukan 9: ")
	printSequence(9, 9)
	fmt.Println()
}
