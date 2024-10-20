package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	baris(n)
	fmt.Println()
}

func baris(bilangan int) {
	for i := bilangan; i >= 1; i-- {
		fmt.Print(i, " ")
	}

	for i := 2; i <= bilangan; i++ {
		fmt.Print(i, " ")
	}
}
