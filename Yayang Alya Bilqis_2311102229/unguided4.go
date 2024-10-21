package main

import (
	"fmt"
	"os"
	"strconv"
)

// Fungsi rekursif untuk mencetak deret angka dari N hingga 1 dan kembali ke N
func cetakDeret(n, current int) {
	if current > n {
		return
	}
	fmt.Print(current, " ")
	if current == 1 {
		for i := 2; i <= n; i++ {
			fmt.Print(i, " ")
		}
		return
	}
	cetakDeret(n, current-1)
	fmt.Print(current, " ")
}

// Fungsi validasi input untuk memastikan input adalah bilangan bulat positif
func validasiInput(input string) (int, error) {
	angka, err := strconv.Atoi(input)
	if err != nil || angka < 1 {
		return 0, fmt.Errorf("input tidak valid, harap masukkan bilangan bulat positif")
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

	fmt.Printf("Deret dari %d hingga 1 dan kembali ke %d: ", n, n)
	cetakDeret(n, n)
	fmt.Println()
}
