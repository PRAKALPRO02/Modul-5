package main

import "fmt"

func cariFaktor(n, i int) {
	if i > n {
		return
	}
	if n % i == 0 {
		fmt.Print(i, " ") // jika n dibagi habis, maka i adalah faktor dari 
	}
	cariFaktor(n, i+1) 
	fmt.Println()
}

func main(){
	var n int 

	fmt.Println("\n*PROGRAM MENAMPILKAN FAKTOR BILANGAN*")
	fmt.Print("\nMasukkan bilangan : ")
	fmt.Scan(&n)

	fmt.Println("\nFaktor dari", n, " adalah:")
	cariFaktor(n, 1) 
	fmt.Println()
}