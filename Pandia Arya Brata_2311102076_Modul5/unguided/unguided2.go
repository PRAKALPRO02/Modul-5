package main

import "fmt"

func main() {
	var tinggi int
	fmt.Scan(&tinggi)
	segitiga(tinggi, 1)
}

func segitiga(n, i int) {
	if i > n {
		return
	} else {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
		segitiga(n, i+1)
	}
}
