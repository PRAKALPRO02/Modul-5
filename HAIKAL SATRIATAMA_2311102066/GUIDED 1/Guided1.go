package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	printDescending(n)
}

func printDescending(n int) {
	if n <= 0 {
		return
	}
	fmt.Println(n)
	printDescending(n - 1)
}
