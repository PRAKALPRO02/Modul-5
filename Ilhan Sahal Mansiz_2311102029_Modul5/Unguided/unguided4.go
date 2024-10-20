package main

import (
    "fmt"
)

func printSequence(N, current int) {
    fmt.Printf("%d ", current)
    if current > 1 {
        printSequence(N, current - 1)
        fmt.Printf("%d ", current)
    }
}

func main() {
    var N int
    fmt.Print("Masukkan bilangan bulat positif N: ")
    fmt.Scan(&N)
    fmt.Printf("Barisan bilangan dari %d hingga 1 dan kembali ke %d:\n", N, N)
    printSequence(N, N)
    fmt.Println()
}