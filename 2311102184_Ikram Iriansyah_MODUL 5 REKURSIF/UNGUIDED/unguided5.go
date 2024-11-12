package main

import "fmt"

func printGanjil(n int, current int) {
	if current > n {
		return
	}
	if current%2 != 0 {
		fmt.Print(current, " ")
	}
	printGanjil(n, current+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif N : ")
	fmt.Scan(&n)
	fmt.Print("Bilangan Ganjil dari 1 -  ", n, " adalah : ")
	printGanjil(n, 1)
	fmt.Println()
}