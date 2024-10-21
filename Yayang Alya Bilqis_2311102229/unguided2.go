package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func cetakBintang(n int, simbol rune, ascending bool) {
	if ascending {
		if n > 0 {
			cetakBintang(n-1, simbol, ascending)
			fmt.Println(string(ulangi(simbol, n)))
		}
	} else {
		if n > 0 {
			fmt.Println(string(ulangi(simbol, n)))
			cetakBintang(n-1, simbol, ascending)
		}
	}
}

func ulangi(karakter rune, jumlah int) []rune {
	var output []rune
	for i := 0; i < jumlah; i++ {
		output = append(output, karakter)
	}
	return output
}

func validasiInput(input string) (int, error) {
	jumlah, err := strconv.Atoi(input)
	if err != nil || jumlah < 1 {
		return 0, fmt.Errorf("input tidak valid")
	}
	return jumlah, nil
}

func pilihSimbol() rune {
	var simbol string
	fmt.Print("Masukkan simbol yang ingin digunakan (default: *): ")
	fmt.Scanln(&simbol)
	if simbol == "" {
		return '*'
	}
	return rune(simbol[0])
}

func main() {
	var pilihan string
	fmt.Println("Pilih mode pola:")
	fmt.Println("1. Pola Meningkat")
	fmt.Println("2. Pola Menurun")
	fmt.Print("Pilihan: ")
	fmt.Scanln(&pilihan)

	fmt.Print("Masukkan jumlah baris: ")
	var input string
	fmt.Scanln(&input)
	baris, err := validasiInput(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	simbol := pilihSimbol()

	switch strings.TrimSpace(pilihan) {
	case "1":
		fmt.Println("Pola Meningkat:")
		cetakBintang(baris, simbol, true)
	case "2":
		fmt.Println("Pola Menurun:")
		cetakBintang(baris, simbol, false)
	default:
		fmt.Println("Pilihan tidak valid.")
		os.Exit(1)
	}
}
