package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(BilBalik(n))
}

func BilBalik(n int) string {
	if n == 1 {
		return "1"
	}
	return fmt.Sprintf("%d %s %d", n, BilBalik(n-1), n)
}
