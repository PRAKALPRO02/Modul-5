package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukan bilangan bulat: ")
	fmt.Scan(&n)
	fmt.Print("Hasil Penjumlahan: ")
	fmt.Println(penjumlahan(n))
}

func penjumlahan(n int) int {
	if n == 1 {
		return 1
	} else {
		return n + penjumlahan(n-1)
	}
}
