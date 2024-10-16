package main

import "fmt"

var reverse bool

func rekursif(n int) {
	if n > 1 {
		fmt.Print(n, " ")
		rekursif(n - 1)
	}
	fmt.Print(n, " ")
}

func main() {
	var a int
	fmt.Scan(&a)
	rekursif(a)
}
