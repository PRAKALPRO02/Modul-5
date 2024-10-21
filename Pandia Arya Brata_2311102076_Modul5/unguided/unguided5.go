package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	bilganjil(num, 1)
}

func bilganjil(n, i int) {
	if i > n {
		fmt.Print(" ")
	} else {
		if i%2 == 1 {
			fmt.Printf("%d ", i)
			bilganjil(n, i+2)
		}
	}
}
