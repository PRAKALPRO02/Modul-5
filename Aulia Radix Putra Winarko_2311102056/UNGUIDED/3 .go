package main

import "fmt"

func faktorial(n, i int, faktors []int) []int {
	if i > n {
		return faktors
	}
	if n%i == 0 {
		faktors = append(faktors, i)
	}
	return faktorial(n, i+1, faktors)
}

func main() {
	var n int
	fmt.Print("Masukan Bilangan Bulat Positif N: ")
	fmt.Scan(&n)
	faktors := faktorial(n, 1, []int{})
	fmt.Println("Faktor dari", n, "Adalah", faktors)
}
