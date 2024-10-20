package main

import (
    "fmt"
)

func findFactors(n int, i int) {
    if i > n {
        return
    }
    if n % i == 0 {
        fmt.Printf("%d ", i)
    }
    findFactors(n, i + 1)
}

func main() {
    var n int
    fmt.Print("Masukkan bilangan bulat positif N: ")
    fmt.Scan(&n)
    fmt.Printf("Faktor dari %d adalah: ", n)
    findFactors(n, 1)
}