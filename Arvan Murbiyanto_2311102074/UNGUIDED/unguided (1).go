package main

import "fmt"

func baris(bilangan int) {
	if bilangan == 1 {
		fmt.Println(1)
	} else {
		fmt.Println(bilangan)
		baris(bilangan - 1)
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai n: ") 
	fmt.Scan(&n)
	
	fmt.Println("Urutan baris:")
	baris(n)
	
	fmt.Println("\nUrutan Fibonacci:")
	for i := 0; i <= n; i++ {
		fmt.Printf("Fibonacci ke-%d: %d\n", i, fibonacci(i))
	}
}
