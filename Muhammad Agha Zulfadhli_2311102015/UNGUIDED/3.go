package main

import "fmt"

func rekursif(n, current int) {
	if n >= current {
		if n%current == 0 {
			fmt.Print(current, " ")
		}
		rekursif(n, current+1)
	}
}

func main() {
	var a int
	fmt.Scan(&a)
	rekursif(a, 1)
}
