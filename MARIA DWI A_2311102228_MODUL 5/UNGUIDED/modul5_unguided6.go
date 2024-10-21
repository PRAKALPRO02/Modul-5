package main

import "fmt"

func pangkat(a, b int)int{
	if (b == 0) {
	return 1
	} else {
		return (a * pangkat(a, b-1))
	}
}

func main(){

	var a, b int
	fmt.Println("\n*PROGRAM MENGHITUNG PANGKAT*")
	fmt.Print("\nMasukkan basis : ")
	fmt.Scan(&a)
	fmt.Print("Masukkan pangkat : ")
	fmt.Scan(&b)

	fmt.Println("\nHasil pangkat : ", pangkat(a, b))
	fmt.Println()
}