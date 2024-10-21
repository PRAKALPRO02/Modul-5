package main

import "fmt"

func main() {
	var bilangan int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bilangan)
	fmt.Print("Faktor dari ", bilangan, " adalah: ")
	Faktor(bilangan, bilangan)
	fmt.Println()
}

func Faktor(bilangan, pembagi int) {
	if pembagi == 1 {
		fmt.Print(1)
		return
	}
	if pembagi == 1 {
		return
	}
	if bilangan%pembagi == 0 {
		fmt.Print(pembagi, " ")
	}
	Faktor(bilangan, pembagi-1)
}
