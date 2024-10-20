package main

import "fmt"

func pangkat(x, y int) int {
	if y == 0 {
		return 1
	} else if y > 0 {
		return x * pangkat(x, y-1)
	} else {
		return 1 / pangkat(x, -y)
	}
}

func main() {
	var x, y int
	fmt.Print("Masukkan bilangan x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan y (pangkat): ")
	fmt.Scan(&y)

	hasil := pangkat(x, y)
	fmt.Printf("%d pangkat %d adalah %d\n", x, y, hasil)
}

// Zahrina Antika Malahati_2311102109
