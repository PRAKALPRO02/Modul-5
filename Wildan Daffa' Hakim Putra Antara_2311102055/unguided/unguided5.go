package main

import "fmt"

func cetakGanjil(n, current int) {
	if current > n {
		return
	} else {
		fmt.Print(current, " ")
		cetakGanjil(n, current+2)
	}
}

func main() {
	var angkaGanjil int
	fmt.Scan(&angkaGanjil)
	cetakGanjil(angkaGanjil, 1)
	fmt.Println()
}
