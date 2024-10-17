package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	barisGanjil(n)
	fmt.Println()
}

func barisGanjil(bilangan int) {
	if bilangan > 0 {
		barisGanjil(bilangan - 1)
		if bilangan%2 != 0 {
			fmt.Print(bilangan, " ")
		}
	}
}
