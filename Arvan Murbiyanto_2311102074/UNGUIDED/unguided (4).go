package main

import "fmt"

func main() {
    var n int
    fmt.Print("Masukkan bilangan: ")
    fmt.Scan(&n)

    fmt.Printf("Barisan bilangan: ")
    printSequence(n)
    fmt.Println()
}

func printSequence(n int) {
    if n == 1 {
        fmt.Print(1, " ")
    } else {
        fmt.Print(n, " ")
        printSequence(n - 1)
        fmt.Print(n, " ")
    }
}
