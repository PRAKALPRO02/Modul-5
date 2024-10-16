package main

import "fmt"

func rekursif(x, y, current int) int {
	hasil := 1
	if current <= y {
		hasil *= (x * (rekursif(x, y, current+1)))
	}
	return hasil
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	fmt.Println(rekursif(x, y, 1))
}
