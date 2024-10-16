package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)
	fmt.Println("Faktor dari", n, "adalah:")
	cetakFaktor(n, 1)
}

func cetakFaktor(n, i int) {
	if i > n {
		return
	}

	if n%i == 0 {
		fmt.Print(i, " ")
	}

	cetakFaktor(n, i+1)
}
