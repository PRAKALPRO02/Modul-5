package main

import "fmt"

func printPattern(n int) {
	for i := 1; i <= n; i++ { 
		for j := 0; j < i; j++ { 
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	var n int
	fmt.Print("Masukkan : ")
	fmt.Scan(&n)
	printPattern(n)
}
