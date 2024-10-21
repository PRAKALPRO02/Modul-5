package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		var result int = fibonacci(n-1) + fibonacci(n-2)
		return result
	}
}

func nthFibonacci(num int) {
	if num == 0 {
		fmt.Printf("%v ", 0)
	} else if num == 1 {
		nthFibonacci(0)
		fmt.Printf("%v ", 1)
	} else {
		nthFibonacci(num - 1)
		fmt.Printf("%v ", fibonacci(num))
	}
}

func main() {
	var num int
	fmt.Scanln(&num)
	nthFibonacci(num)
}
