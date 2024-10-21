package main

import "fmt"

func faktor(n, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Printf("%d ", i)
	}
	faktor(n, i+1)
}

func main() {
	var n int
	fmt.Scanln(&n)
	faktor(n, 1)
}
