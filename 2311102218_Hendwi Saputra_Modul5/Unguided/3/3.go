package main

import "fmt"

func cariFaktor(bilangan, pembagi int) {
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
	cariFaktor(bilangan, pembagi-1)
}

func main() {
	var bilangan int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bilangan)
	fmt.Print("Faktor dari ", bilangan, " adalah: ")
	cariFaktor(bilangan, bilangan)
	fmt.Println()
}
