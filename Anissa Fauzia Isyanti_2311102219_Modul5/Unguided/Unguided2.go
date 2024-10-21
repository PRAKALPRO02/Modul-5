package main

import "fmt"

func cetak_bintang219(n int) {
	if n == 0 {
		return
	} else {
		fmt.Print("* ")
		cetak_bintang219(n - 1)
	}
}

func pola(n, i int) {
	if i > n {
		return
	} else {
		cetak_bintang219(i)
		fmt.Println()
		pola(n, i+1)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	pola(n, 1)
}
