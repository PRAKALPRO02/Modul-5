package main

import "fmt"

func bilGanjil(n int) {
	for i := 1; i <= n; i += 2 { 
		fmt.Print(i, " ")
	}
}

func main() {
	var n int
	fmt.Print("Masukkan: ")
	fmt.Scan(&n)

	bilGanjil(n) 
	fmt.Println()
}
