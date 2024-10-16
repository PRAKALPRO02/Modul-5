package main

import "fmt"

func cetak(n int) {
	if n > 0 {
		fmt.Print("*")
		cetak(n - 1)
	}
}

func rekursif(n int) {
	if n == 1 {
		fmt.Println("*")
	} else {
		rekursif(n - 1)
		cetak(n)
		fmt.Println()
	}
}

func main() {
	var a int
	fmt.Scan(&a)
	rekursif(a)
}
