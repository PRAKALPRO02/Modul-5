package main

import "fmt"

func mencetakBilangan(n int) {
	if n > 1 {
		fmt.Print(n, " ")
		mencetakBilangan(n - 1)
		fmt.Print(n, " ")
	} else if n == 1 {
		fmt.Print(n, " ")
	}
}

func main() {
	var n int
	for {
		fmt.Print("Masukan bilangan: ")
		fmt.Scan(&n)
		if n > 0 {
			break
		}
	}
	mencetakBilangan(n)
	fmt.Println()
}
