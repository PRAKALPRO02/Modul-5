package main

import (
	"fmt"
)

func cetakBintang(n int) {
	if n == 0 {
		return
	}
	fmt.Print("* ")
	cetakBintang(n - 1)
}

func polaBintang(n, i int) {
	if i > n {
		return
	}
	cetakBintang(i)
	fmt.Println()
	polaBintang(n, i+1)
}

func main() {
	var n int
	fmt.Scan(&n)
  polaBintang(n, 1)
}