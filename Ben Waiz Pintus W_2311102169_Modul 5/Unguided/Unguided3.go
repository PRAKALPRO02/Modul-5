//Ben Waiz Pintus W
//2311102169

package main

import "fmt"

func cetakFaktor(n, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Printf("%d ", i)
	}
	cetakFaktor(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)

	fmt.Printf("Nilai faktor dari %d adalah: ", n)
	cetakFaktor(n, 1)
	fmt.Println()
}
