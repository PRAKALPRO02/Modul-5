package main

import "fmt"

func faktor(n, i int) {
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	if n == i {
		return
	} else {
		faktor(n, i+1)
	}
}

func main() {
	var faktorN int
	fmt.Scan(&faktorN)
	faktor(faktorN, 1)
}
