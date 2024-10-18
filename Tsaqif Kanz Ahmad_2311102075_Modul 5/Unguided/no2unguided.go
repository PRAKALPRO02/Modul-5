package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scanln(&n)

	printStars(n)
}

func printStars(n int) {
	if n == 0 {
		return
	}
	printStars(n - 1)
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}
