package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	barisGanjil(n)
	fmt.Println()
}

func barisGanjil(bilangan int) {
	for i := 1; i <= bilangan; i++ {
		if i%2 != 0 {
			fmt.Print(i, " ")
		}
	}
}
