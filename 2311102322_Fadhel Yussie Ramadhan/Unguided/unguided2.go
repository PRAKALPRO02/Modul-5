package main

import "fmt"

func main () {
	var a int
	fmt.Print ("Masukan nilai: ")
	fmt.Scan(&a)
	pola(a, 1)

}

func bintang (B int) {
	if B == 0 {
		return
	}
	fmt.Print ("* ")
	bintang (B - 1)
}

func pola (n, c int) {
	if c > n {
		return
	}
	bintang (c)
	fmt.Println()
	pola (n, c+1)
}