package main

import "fmt"

func barisBilangan(num int) {
	if num < 1 {
		fmt.Println("Bilangan harus positif")
	} else if num == 1 {
		fmt.Print("1 ")
	} else {
		fmt.Printf("%v ", num)
		barisBilangan(num - 1)
		fmt.Printf("%v ", num)
	}
}

func main() {
	var num int
	fmt.Scanln(&num)
	barisBilangan(num)
}
