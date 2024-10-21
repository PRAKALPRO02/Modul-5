package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan bilangan (x): ")
	fmt.Scan(&x)
	fmt.Print("Masukkan pangkat (y): ")
	fmt.Scan(&y)

	hasil := pangkat(x, y)
	fmt.Println("Hasil:", hasil)
}

func pangkat(x, y int) int {
	if y == 0 {
		return 1
	}
	return x * pangkat(x, y-1)
}
