package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println("Faktor dari", n, "adalah:")
	cetakFaktor(n)
}

func cetakFaktor(n int) {
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
