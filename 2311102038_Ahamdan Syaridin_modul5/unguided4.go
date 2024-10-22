package main

import "fmt"

func printSequence(n int) {
	
	for i := n; i >= 1; i-- {
		fmt.Print(i, " ")
	}
	
	for i := 2; i <= n; i++ {
		fmt.Print(i, " ")
	}
}

func main() {
	var n int
	fmt.Print("Masukkan : ")
	fmt.Scan(&n)
	
	printSequence(n) 
	fmt.Println()
}
