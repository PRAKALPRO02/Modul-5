package main

import "fmt"

func printFactors(n int) {
	for i := 1; i <= n; i++ { 
		if n%i == 0 { 
			fmt.Print(i, " ") 
		}
	}
}

func main() {
	var n int
	fmt.Print("Masukkan : ")
	fmt.Scan(&n)
	fmt.Print("Faktor dari ", n, " adalah: ")
	printFactors(n) 
	fmt.Println()
}
