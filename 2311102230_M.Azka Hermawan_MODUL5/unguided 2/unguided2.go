package main

import "fmt"

func main() {
	var N int
	fmt.Print("Masukan jumlah baris yang diinginkan: ")
	fmt.Scan(&N)
	pola(N, 1)
}

func cetakBintang(n, i int) {
	if i == 0 {
		return
	}

	fmt.Print("*")
	cetakBintang(n, i-1)
}

func pola(n, baris int) {
	if baris > n {
		return
	}
	cetakBintang(n, baris)
	fmt.Println()
	pola(n, baris+1)
}
