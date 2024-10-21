package main

import "fmt"

func segitiga(i, n int) {
	if i > n {
		return
	}
	for j := 0; j < i; j++ {
		fmt.Print("*")
	}
	fmt.Println()
	segitiga(i+1, n)
}

func main() {
	var n int
	fmt.Print("masukan tinggi segitiga : ")
	fmt.Scan(&n)
	segitiga(1, n)
}
