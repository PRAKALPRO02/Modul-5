package main

import "fmt"

func printDescending(n int) {
	if n == 0 {
		return
	}
	fmt.Print(n, " ")
	printDescending(n - 1)
}

func printAscending(n, current int) {
	if current > n {
		return
	}
	fmt.Print(current, " ")
	printAscending(n, current+1)
}

func main() {
	var n int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	printDescending(n)

	printAscending(n, 2)
}
