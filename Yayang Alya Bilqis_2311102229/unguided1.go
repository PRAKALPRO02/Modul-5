package main

import "fmt"

var memo = map[int]int{0: 0, 1: 1}

func fibonacci(n int) int {
	if val, ok := memo[n]; ok {
		return val
	}
	if n <= 1 {
		return n
	}
	result := fibonacci(n-1) + fibonacci(n-2)
	memo[n] = result
	return result
}

func fibonacciIterative(n int) int {
	if n <= 1 {
		return n
	}
	prev, current := 0, 1
	for i := 2; i <= n; i++ {
		next := prev + current
		prev, current = current, next
	}
	return current
}

func main() {
	fmt.Println("Menggunakan rekursi dengan memoization:")
	for i := 0; i <= 10; i++ {
		fmt.Printf("Fibonacci(%d) = %d\n", i, fibonacci(i))
	}

	fmt.Println("\nMenggunakan iterasi:")
	for i := 0; i <= 10; i++ {
		fmt.Printf("Fibonacci(%d) = %d\n", i, fibonacciIterative(i))
	}
}
