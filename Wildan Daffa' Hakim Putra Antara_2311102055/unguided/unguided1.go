package main

import (
	"fmt"
	"strconv"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	var inp int
	var bilke, fibo string
	fmt.Scan(&inp)
	for i := 1; i <= inp; i++ {
		bilke += "\t" + strconv.Itoa(i)
		fibo += "\t" + strconv.Itoa(fibonacci(i))
	}
	fmt.Printf("n: %s \nSn: %s \n", bilke, fibo)
}
