package main

import (
	"fmt"
	"os"
	"strconv"
)

func cetakBilanganGanjil(n, current int) {
	if current > n {
		return
	}
	fmt.Print(current, " ")
	cetakBilanganGanjil(n, current+2)
}

func validasiInput(input string) (int, error) {
	angka, err := strconv.Atoi(input)
	if err != nil || angka < 1 {
		return 0, fmt.Errorf("Input tidak valid, harap masukkan bilangan bulat positif")
	}
	return angka, nil
}

func main() {
	fmt.Print("Masukkan bilangan bulat positif N: ")
	var input string
	fmt.Scanln(&input)
	n, err := validasiInput(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Deret bilangan ganjil dari 1 hingga %d adalah: ", n)
	cetakBilanganGanjil(n, 1)
	fmt.Println()
}
