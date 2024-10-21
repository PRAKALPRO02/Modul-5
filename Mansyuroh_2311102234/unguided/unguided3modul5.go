package main

import "fmt"

func faktorBilangan234(bilangan234, i234 int) {
	if i234 > bilangan234 {
		return
	}
	if bilangan234%i234 == 0 {
		fmt.Print(i234, " ")
	}
	faktorBilangan234(bilangan234, i234+1)
}

func main() {
	var inputBilangan234 int
	fmt.Print("\nMasukkan bilangan : ")
	fmt.Scan(&inputBilangan234)
	fmt.Println("\nFaktor dari", inputBilangan234, " adalah:")
	faktorBilangan234(inputBilangan234, 1)
	fmt.Println()
}
