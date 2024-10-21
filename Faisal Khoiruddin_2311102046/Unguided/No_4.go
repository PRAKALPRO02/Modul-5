package main

import "fmt"

func descending(n int) {
	if n < 1 {
		return
	}
	fmt.Printf("%d ", n)
	descending(n - 1)
}

func ascending(n int, current int) {
	if current > n {
		return
	}
	fmt.Printf("%d ", current)
	ascending(n, current+1)
}

func main() {
	var n int
	fmt.Scanln(&n)
	descending(n)
	ascending(n, 2)
}
