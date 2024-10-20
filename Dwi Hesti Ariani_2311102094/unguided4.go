package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scanln(&n)

	fmt.Println("Keluaran:", generateSequence(n))
}

func generateSequence(n int) string {
	if n == 1 {
		return "1"
	}

	return fmt.Sprintf("%d %s %d", n, generateSequence(n-1), n)
}
