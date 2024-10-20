package main

import "fmt"

func cetak_BilanganGanjil(n int) {
	if n <= 0 {
		return
	}
	cetak_BilanganGanjil(n - 1)
	if n%2 != 0 {
		fmt.Println(n)
	}
}

func main() {
	var n int

	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&n)

	fmt.Println("Barisan bilangan ganjil: ")
	cetak_BilanganGanjil(n)

}
