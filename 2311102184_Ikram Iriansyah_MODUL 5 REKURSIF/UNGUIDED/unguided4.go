package main

import "fmt"


func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scanln(&n)
	printRecursive(n)
}

func printRecursive(n int) {
	if n == 1 {
		fmt.Print(n, " ")
		return
	}

	fmt.Print(n, " ")

	printRecursive(n - 1)
	
	fmt.Print(n, " ")
}