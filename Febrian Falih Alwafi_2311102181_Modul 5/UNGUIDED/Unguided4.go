package main

import "fmt"

func main() {
    var n int
    fmt.Print("Masukkan bilangan bulat positif: ")
    fmt.Scanln(&n)
    fmt.Print("Barisan bilangan : ")
    printSequence(n)
    
}

func printSequence(n int) {
    if n == 1 {
        fmt.Print(n, " ")
        return
    } //2311102181
    fmt.Print(n, " ")
    printSequence(n - 1)
    fmt.Print(n, " ")
}


