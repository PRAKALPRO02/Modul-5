package main

import "fmt"

func baris(n, i int) {
	if i > 2*n {
		return
	} else {
		if i < n {
			fmt.Print(n-i, " ")
		} else if i == n || i == n+1 {
			fmt.Print()
		} else {
			fmt.Print(i-n, " ")
		}
		baris(n, i+1)
	}
}
func main() {
	var n int
	fmt.Scan(&n)
	baris(n, 0)
}
