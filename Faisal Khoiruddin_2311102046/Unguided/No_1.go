package main

import (
	"fmt"
	"strings"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)

	fibs := make([]int, n+1)
	for i := 0; i <= n; i++ {
		fibs[i] = fibonacci(i)
	}

	colWidth := 6
	tableWidth := (n + 2) * colWidth
	line := strings.Repeat("-", tableWidth)

	fmt.Println(line)
	fmt.Printf("| %-3s ", "n ")
	for i := 0; i <= n; i++ {
		fmt.Printf("| %-3d ", i)
	}
	fmt.Println("|")
	fmt.Println(line)
	fmt.Printf("| %-3s ", "Sn")
	for i := 0; i <= n; i++ {
		fmt.Printf("| %-3d ", fibs[i])
	}
	fmt.Println("|")
	fmt.Println(line)
}
