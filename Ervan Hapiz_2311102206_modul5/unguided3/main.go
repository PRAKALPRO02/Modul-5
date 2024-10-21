package main

import "fmt"

func faktorial(n, i int) {
	if i > n {
		return
	} else {
		if n%i == 0 {
			fmt.Print(i, " ")
		}
		faktorial(n, i+1)
	}
}
func main() {
	var n int
	fmt.Scan(&n)
	faktorial(n, 1)
}
