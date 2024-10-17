package main

import "fmt"

func pangkat(x, y int) int {
	if y == 0 {
		return 1
	} else {
		return x * pangkat(x, y-1)
	}
}

func main() {
	var base, exponent int

	fmt.Print("Masukkan bilangan pokok (x) : ")
	fmt.Scan(&base)
	fmt.Print("Masukkan pangkat (y) : ")
	fmt.Scan(&exponent)

	result := pangkat(base, exponent)
	fmt.Printf("%d pangkat %d adalah : %d\n", base, exponent, result)
}
