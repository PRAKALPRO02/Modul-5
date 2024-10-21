package main

import "fmt"

func power(x, y int) int {
	result := 1
	for i := 0; i < y; i++ {
		result *= x
	}
	return result
}

func main() {
	fmt.Println("Masukkan 2 bilangan bulat (x, y):")
	var x, y int
	fmt.Scanln(&x, &y)

	result := power(x, y)
	fmt.Printf("%d pangkat %d = %d\n", x, y, result)
}
