package main

import "fmt"

func barisTurun(n int) {
	if n <= 0 {
		return
	}
	fmt.Print(n, " ")
	barisTurun(n - 1)
}

func barisNaik(n, awal int) {
	if n > awal {
		return
	}
	fmt.Print(n, " ")
	barisNaik(n+1, awal)
}

func main() {
	var o int
	fmt.Print("Masukan bilangan bulat positif: ")
	fmt.Scan(&o)

	barisTurun(o)

	if o >= 1 {
		barisNaik(2, o)
	}

	fmt.Println()
}
