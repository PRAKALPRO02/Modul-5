package main

import "fmt"

func main (){
	var n  int
	fmt.Print("Masukan Bilangan: ")
	fmt.Scan (&n)
	baris (n)
}

func baris (bilangan int) {
	if bilangan < 1 {
		return
	}
	fmt.Print (bilangan," " )
	baris(bilangan - 1)
	if bilangan > 1 {
		fmt.Print (bilangan," " )
	}
}