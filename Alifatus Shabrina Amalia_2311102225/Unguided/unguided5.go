package main

import "fmt"

func printOddNumbers(n, current int) {
	if current > n {
		return
	}

	if current%2 != 0 {
		fmt.Print(current, " ")
	}

	printOddNumbers(n, current+1)
}

func main() {
	var n int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	fmt.Println("Bilangan ganjil dari 1 hingga", n, "adalah:")
	printOddNumbers(n, 1)
}
