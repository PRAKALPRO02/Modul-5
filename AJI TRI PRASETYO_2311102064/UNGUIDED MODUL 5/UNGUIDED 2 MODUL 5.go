package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)
	fmt.Println(printbintang(n))
}

func printbintang(n int) string {
	if n == 1 {
		return "*"
	}
	return strings.Join([]string{printbintang(n - 1), cetakBintang(n)}, "\n")
}

func cetakBintang(n int) string {
	return strings.Repeat("*", n)
}
