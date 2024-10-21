package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)

	fmt.Println("Faktor-faktor dari", bilangan, " adalah ")
	faktor(bilangan, 1)
}

func faktor(n, f int) {
	if f > n {
		return
	}
	if n%f == 0 {
		fmt.Println(f)
	}
	faktor(n, f+1)
}
