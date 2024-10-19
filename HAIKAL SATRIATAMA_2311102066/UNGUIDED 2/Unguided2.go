package main

import "fmt"

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

func main() {
	var N int
	for {
		fmt.Print("Masukkan angka: ")
		fmt.Scan(&N)
		if N == 0 {
			fmt.Println("Program selesai.")
			break
		}
		printStars(N)
	}
}
