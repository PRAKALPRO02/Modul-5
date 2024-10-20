package main

import (
    "fmt"
)

func power(x, y int) int {
    if y == 0 {
        return 1
    }
    return x * power(x, y-1)
}

func main() {
    var x, y int
    fmt.Print("Masukkan nilai x dan y: ")
    fmt.Scan(&x, &y)

    hasil := power(x, y)
    fmt.Printf("Hasil dari %d dipangkatkan %d adalah: %d\n", x, y, hasil)
}