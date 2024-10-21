package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukan Angka: ")
	fmt.Scan(&n)
	fmt.Println("Hasilnya")
	baris(n)
}

func baris(bilangan int) {
	if bilangan == 1 {
		fmt.Println(1)
	} else {
		fmt.Println(bilangan)
		baris(bilangan - 1)
	}
}
