package main

import "fmt"

// Fungsi rekursif untuk mencetak bintang 
func printStars(n int) {
	if n == 0 {
		return
	}
	printStars(n - 1)
	fmt.Println(string(make([]rune, n, n))) // Mencetak n bintang
}

// Fungsi rekursif untuk mencetak pola hingga N baris
func printPattern(n, i int) {
	if i > n {
		return
	}
	for j := 0; j < i; j++ {
		fmt.Print("*")
	}
	fmt.Println()
	printPattern(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan : ")
	fmt.Scan(&n)
	printPattern(n, 1) // Memulai dari baris pertama
}
