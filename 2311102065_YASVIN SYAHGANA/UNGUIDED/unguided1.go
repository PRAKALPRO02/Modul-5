package main

import "fmt"

func fibonaci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonaci(n-1) + fibonaci(n-2)
}

func main() {
	//var n int
	fmt.Print("n : ")
	for i := 0; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Print("\nSn: ")
	for j := 0; j <= 10; j++ {
		fmt.Print(fibonaci(j), " ")
	}
}
