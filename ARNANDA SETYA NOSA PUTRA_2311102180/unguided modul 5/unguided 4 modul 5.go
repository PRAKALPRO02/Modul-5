package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	baris(n)
	fmt.Println()
}

func baris(bilangan int) {
	if bilangan == 1 {
		fmt.Print(1, " ")
	} else {
		fmt.Print(bilangan, " ")
		baris(bilangan - 1)
		fmt.Print(bilangan, " ")
	}
}
