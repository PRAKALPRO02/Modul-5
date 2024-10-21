package main

// 2311102002

import "fmt"

// Fungsi rekursif untuk menghitung deret Fibonacci
func fibonacci(n int) int {
	if n <= 1 {
		return n // Jika nilai n adalah 0 atau 1, maka langsung kembalikan nilainya
	}
	return fibonacci(n-1) + fibonacci(n-2) // Rekursi untuk menghitung nilai suku sebelumnya
}

func main() {
	fmt.Println("Deret Fibonacci hingga suku ke-10:")
	for i := 0; i <= 10; i++ { // Langsung gunakan angka 10 di sini
		fmt.Printf("Suku ke-%d: %d\n", i, fibonacci(i))
	}
}
