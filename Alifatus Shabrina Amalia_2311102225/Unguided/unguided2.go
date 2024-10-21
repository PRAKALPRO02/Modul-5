package main

import "fmt"

func printStars(n int) {
	if n == 0 {
		return
	}
	fmt.Print("*")
	printStars(n - 1)
}

func printPattern(n int, current int) {
	if current > n {
		return
	}
	printStars(current)
	fmt.Println()
	printPattern(n, current+1)
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&n)

	fmt.Println("Pola bintang:")
	printPattern(n, 1)
}
