package main

import (
	"fmt"
	"os"
	"strconv"
)

func faktorRekursif(n, divisor int) {
	if divisor > n {
		return
	}
	if n%divisor == 0 {
		fmt.Print(divisor, " ")
	}
	faktorRekursif(n, divisor+1)
}

func validasiInput(input string) (int, error) {
	jumlah, err := strconv.Atoi(input)
	if err != nil || jumlah < 1 {
		return 0, fmt.Errorf("input tidak valid")
	}
	return jumlah, nil
}

func main() {
	fmt.Print("Masukkan bilangan bulat positif: ")
	var input string
	fmt.Scanln(&input)
	n, err := validasiInput(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Faktor dari %d adalah: ", n)
	faktorRekursif(n, 1)
	fmt.Println()
}
