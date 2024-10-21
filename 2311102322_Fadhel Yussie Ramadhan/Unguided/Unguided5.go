package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)
	barisGanjil(num)
}

func barisGanjil(num int) {
	if num <= 1 {
		fmt.Print("1 ")
	} else if num%2 != 0 {
		barisGanjil(num - 1)
		fmt.Printf("%v ", num)
	} else {
		barisGanjil(num - 1)
	}
}