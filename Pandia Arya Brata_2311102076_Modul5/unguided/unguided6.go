package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	hasil := pangkat(x, y)
	fmt.Print(hasil)
}

func pangkat(x, y int) int {
	if y == 0 {
		return 1
	} else {
		return x * pangkat(x, y-1)
	}
}
