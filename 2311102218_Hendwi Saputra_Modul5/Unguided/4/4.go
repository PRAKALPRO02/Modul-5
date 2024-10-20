package main

import "fmt"

func cetakBaris(n int) {
	if n > 1 {
		fmt.Print(n, " ")
		cetakBaris(n - 1)
		fmt.Print(n, " ")
	} else if n == 1 {
		fmt.Print(n, " ")
	}
}

func main() {
	var n int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)
	cetakBaris(n)
}
