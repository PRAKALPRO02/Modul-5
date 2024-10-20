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

func polaBintang(n int) {
	for i := 1; i <= n; i++ {
		cetakBintang(i)
		fmt.Println()
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	polaBintang(n)
}

