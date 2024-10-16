package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)
	fmt.Println(printbintang(n))
}

func printbintang(n int) string {
	if n == 1 {
		return "*"
	} else {
		return printbintang(n-1) + "\n" + cetakBintang(n)
	}
}

func cetakBintang(n int) string {
	bintang := ""
	for i := 0; i < n; i++ {
		bintang += "*"
	}
	return bintang
}
