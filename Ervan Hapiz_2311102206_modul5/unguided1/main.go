package main

import "fmt"

func fibonaci(n int) int {
	if n <= 1 {
		return 1
	} else {
		return fibonaci(n-1) + fibonaci(n-2)
	}
}

func main() {
	var n int
	fmt.Print("masukan nilai n :")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Print(fibonaci(i), " ")
	}

}
