package main

import (
    "fmt"
)

func printOdds(current, N int) {
    if current > N {
        return
    }
    if current % 2 != 0 {
        fmt.Printf("%d ", current)
    }
    printOdds(current + 1, N)
}

func main() {
    var N int
    fmt.Print("Masukkan bilangan bulat positif N: ")
    fmt.Scan(&N)
    fmt.Printf("Bilangan ganjil dari 1 hingga %d adalah: ", N)
    printOdds(1, N)
    fmt.Println()
}