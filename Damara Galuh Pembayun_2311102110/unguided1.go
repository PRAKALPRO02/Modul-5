package main

import "fmt"

// fibonacci adalah fungsi rekursif untuk menghitung nilai suku ke-n dari deret Fibonacci
func fibonacci(n int) int {
        // Kasus dasar:
        // Jika n adalah 0 atau 1, nilai Fibonacci-nya adalah n itu sendiri
        if n <= 1 {
                return n
        }

        // Kasus rekursif:
        // Nilai Fibonacci ke-n adalah penjumlahan nilai Fibonacci ke-(n-1) dan ke-(n-2)
        return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
        var n int

        fmt.Print("Masukkan nilai n: ")
        fmt.Scan(&n)

        // Panggil fungsi fibonacci untuk menghitung nilai suku ke-n
        result := fibonacci(n)

        fmt.Printf("Nilai suku ke-%d dari deret Fibonacci adalah: %d\n", n, result)
}