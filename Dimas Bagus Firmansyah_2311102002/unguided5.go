package main

import "fmt"

// Fungsi rekursif untuk menampilkan bilangan ganjil dari 1 hingga N
func bilGanjil(n, current int) {
	if current > n {
		return
	}
	fmt.Print(current, " ")
	bilGanjil(n, current+2)
}

func main() {
	var n int
	fmt.Print("Masukkan: ")
	fmt.Scan(&n)

	bilGanjil(n, 1)
	fmt.Println() // 2311102002
}
