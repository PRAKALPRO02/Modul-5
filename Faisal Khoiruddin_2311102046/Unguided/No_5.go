package main

import "fmt"

func bil_ganjil(n, current int) {
	if current > n {
		return
	}
	if current%2 != 0 {
		fmt.Printf("%d ", current)
	}
	bil_ganjil(n, current+1)
}

func main() {
	var n int
	fmt.Scanln(&n)
	bil_ganjil(n, 1)
}
