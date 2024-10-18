package main

import "fmt"

func pangkat(n,m int) {
	val :=n
	for i:=1; i <m; i++ {
		val = val * n
	}
	fmt.Print(val)
}
func main() {
	var n,m int
	fmt.Scan(&n,&m)
	pangkat(n,m)
}