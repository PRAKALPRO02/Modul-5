package main

import "fmt"

func pangkat(x, y int) int {
	result := 1
	for i := 0; i < y; i++ { 
		result *= x 
	}
	return result
}

func main() {
	var x, y int
	fmt.Print("Masukkan x dan y: ")
	fmt.Scan(&x, &y)

	fmt.Println(pangkat(x, y)) 
}
