package main

import "fmt"

func main() {
	var num, power_of int
	fmt.Scanln(&num, &power_of)
	fmt.Println(power(num, power_of))
}

func power(num int, power_of int) int {
	if power_of <= 1 {
		return num * power_of
	} else {
		return num * power(num, power_of-1)
	}
}