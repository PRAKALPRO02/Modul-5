package main

import "fmt"

func power(x, y int) int {
	if y == 0 {
		return 1
	}
	return x * power(x, y-1)
}

func main() {
	fmt.Println("Masukan 2 bilangan bulat (x, y):")
	var x, y int
	fmt.Scanln(&x, &y)

	result := power(x, y)
	fmt.Printf("%d pangkat %d = %d\n", x, y, result)
}
