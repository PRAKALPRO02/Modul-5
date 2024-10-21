package main

import "fmt"

func factorial(num int, i int) {
	if i >= num {
		fmt.Printf("%v", num)
	} else {
		if num%i == 0 {
			fmt.Printf("%v ", i)
		}
		factorial(num, i+1)
	}
}

func main() {
	var num int
	fmt.Scanln(&num)
	factorial(num, 1)
}
