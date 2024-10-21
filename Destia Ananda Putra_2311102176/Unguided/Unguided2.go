package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)
	cetakBintang(n, 1)
}

func cetakBintang(total, current int) {
	if current > total {
		return
	}
	for i := 0; i < current; i++ {
		fmt.Print("*")
	}
	fmt.Println()
	cetakBintang(total, current+1)

	return
}
