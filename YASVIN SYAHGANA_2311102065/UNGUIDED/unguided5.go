package main

import "fmt"

func ganjil(i, n int) {
	if i > n {
		return
	}
	if i%2 != 0 {
		fmt.Print(i, " ")
	}
	ganjil(i+1, n)
}

func main() {
	var n int
	fmt.Print("Input (n): ")
	fmt.Scan(&n)

	ganjil(0, n)
}
