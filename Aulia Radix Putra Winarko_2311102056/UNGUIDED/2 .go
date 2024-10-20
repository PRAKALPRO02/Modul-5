package main

import "fmt"

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

func main () {
	var a int
	fmt.Print ("Masukan nilai N: ")
	fmt.Scan(&a)
	pola(a, 1)

}