package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	baris_219(n, 1)
}

func baris_219(n, bilangan int) {
	if n >= bilangan {
		if bilangan%2 != 0 {
			fmt.Print(bilangan, " ")
		}
		baris_219(n, bilangan+1)
	}
}
