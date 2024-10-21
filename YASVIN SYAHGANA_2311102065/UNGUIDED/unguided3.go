package main

import "fmt"

func faktorisasi(n int, f int) {
	if n < f {
		return
	}
	if n%f == 0 {
		fmt.Print(f, " ")
	}
	faktorisasi(n, f+1)
}

func main() {
	f := 1
	var n int
	fmt.Print("input(n): ")
	fmt.Scanln(&n)

	faktorisasi(n, f)

}
