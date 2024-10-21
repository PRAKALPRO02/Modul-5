package main

import (
	"fmt"
)

func ganjil(n, i int) {
	if i > n {
		return
	} else {
		if i%2 == 1 {
			fmt.Print(i, " ")
		}
		ganjil(n, i+1)
	}
}
func main() {
	var n int

	fmt.Print("masukan bilangan N : ")
	fmt.Scan(&n)
	ganjil(n, 1)

}
