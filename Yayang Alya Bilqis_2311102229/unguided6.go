package main

import (
	"fmt"
)

7func power(base, exponent int) int {
	if exponent == 0 {
		return 1
	} else if exponent < 0 {
		return 1 / power(base, -exponent)
	} else {
		return base * power(base, exponent-1)
	}
}

func powerIterative(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}

func printResult(base, exponent, result int) {
	fmt.Printf("%d pangkat %d = %d\n", base, exponent, result)
}

func main() {
	var base, exponent int

	fmt.Print("Masukkan bilangan pokok: ")
	fmt.Scan(&base)

	fmt.Print("Masukkan pangkat: ")
	fmt.Scan(&exponent)

	recursiveResult := power(base, exponent)

	iterativeResult := powerIterative(base, exponent)

	if recursiveResult == iterativeResult {
		fmt.Println("Hasil perhitungan rekursif dan iteratif sama.")
	} else {
		fmt.Println("Terdapat kesalahan dalam perhitungan.")
	}

	printResult(base, exponent, recursiveResult)
}
