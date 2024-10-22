package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(Penjumlahan(n))
}
     func Penjumlahan(n int) int{
		if n == 1{
			return 1
		} else {
			return n + Penjumlahan(n - 1)
		}
	 }