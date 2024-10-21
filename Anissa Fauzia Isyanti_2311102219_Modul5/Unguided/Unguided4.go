package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	barisan_219(n)
	fmt.Print()
}

func barisan_219(n int) {
	if n < 1 {
		return
	}
	fmt.Print(n, " ")
	barisan_219(n - 1)
	if n > 1 {
		fmt.Print(n, " ")
	}
}
