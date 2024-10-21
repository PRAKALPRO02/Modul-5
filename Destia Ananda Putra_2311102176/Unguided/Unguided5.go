package main

import "fmt"

func cetakGanjil(n, current int) {
	if current > n {
		return
	}
	fmt.Print(current, " ")
	cetakGanjil(n, current+2)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&n)
	cetakGanjil(n, 1)
}
