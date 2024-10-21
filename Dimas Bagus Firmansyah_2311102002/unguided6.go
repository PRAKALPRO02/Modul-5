package main

import "fmt"

// Fungsi rekursif untuk menghitung x pangkat y
func pangkat(x, y int) int {
	if y == 0 {
		return 1
	}
	return x * pangkat(x, y-1)
}

func main() {
	var x, y int
	fmt.Print("Masukkan x dan y: ")
	fmt.Scan(&x, &y) // 2311102002

	fmt.Println(pangkat(x, y))
}
