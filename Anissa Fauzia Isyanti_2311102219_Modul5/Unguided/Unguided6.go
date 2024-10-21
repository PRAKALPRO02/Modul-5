package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	fmt.Print(pangkat(x, y))
}

func pangkat(x, y int) int {
	if y == 0 {
		return 1
	} else {
		return x * pangkat(x, y-1)
	}
}
