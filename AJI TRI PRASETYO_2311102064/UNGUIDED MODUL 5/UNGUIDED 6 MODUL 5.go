package main

import "fmt"

func main() {
	var base, exponent int
	fmt.Scan(&base)
	fmt.Scan(&exponent)

	result := calculatePower(base, exponent)
	fmt.Println(result)
}

func calculatePower(base, exponent int) int {
	if exponent == 0 {
		return 1
	}
	return base * calculatePower(base, exponent-1)
}
