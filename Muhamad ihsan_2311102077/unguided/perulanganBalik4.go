package main

import "fmt"

func loop(n int) {
	for i:=n; i >= 1; i-- {
		fmt.Print(i," ")
		if i == 1 {
			for i = 2; i <= n; i++ {
				fmt.Print(i," ")
			}
			break
		}
	}
}
func main() {
	var n int
	fmt.Scan(&n)
	loop(n)
}