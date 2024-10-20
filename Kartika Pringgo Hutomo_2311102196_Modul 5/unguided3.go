package main

import "fmt"

func rekursif(n int) {
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Print(i, " ")
		}
	}
}

func main() {
	var a int
	fmt.Scan(&a)
	rekursif(a)
}
