package main

import "fmt"

func main() {
	var n int
	//fahrur059
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scanln(&n)

	printStars(n)
}

func printStars(n int) {
	if n == 0 {
		return
	}
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
	printStars(n - 1)
}
