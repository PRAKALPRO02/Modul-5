package main

import "fmt"

func printFactors(bilangan, pembagian int) {
	if pembagian > bilangan {
		return
	}
	if bilangan%pembagian == 0 {
		fmt.Print(pembagian, " ")
	}
	printFactors(bilangan, pembagian+1)
}

func main() {
	var bilangan int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bilangan)

	fmt.Printf("Faktor dari %d adalah: ", bilangan)
	printFactors(bilangan, 1)
	fmt.Println()
}
