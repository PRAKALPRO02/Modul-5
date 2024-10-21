package main

import "fmt"

func main() {
	fmt.Print("Masukkan bilangan bulat positif: ")
	var n int
	fmt.Scanln(&n)
	//fahrur059
	fmt.Println("Faktor dari", n, "adalah:", faktor(n))
}

func faktor(n int) []int {
	var faktor []int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			faktor = append(faktor, i)
		}
	}
	return faktor
}
