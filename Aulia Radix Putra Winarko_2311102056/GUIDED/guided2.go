package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(penjualan(n))
}

func penjualan(n int) int {
	if n == 1 {
		return 1
	} else {
		return n + penjualan(n - 1)
	}
}