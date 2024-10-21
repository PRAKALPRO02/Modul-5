package main

import "fmt"

// Fungsi menampilkan bilangan n hingga 1 dan kembali
func printSequence(n, current int) {
	if current == 0 {
		return
	}
	fmt.Print(current, " ")
	printSequence(n, current-1)
	if current != n {
		fmt.Print(current, " ")
	}
}

func main() {
	var n int
	fmt.Print("Masukkan : ")
	fmt.Scan(&n)
	
	printSequence(n, n)
	fmt.Println() // 2311102002
}
