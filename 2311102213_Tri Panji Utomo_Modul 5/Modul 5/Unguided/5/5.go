package main

import "fmt"

func main() {
    var o int

    fmt.Print("Masukkan bilangan bulat positif: ")
    fmt.Scan(&o)

    fmt.Print("Bilangan ganjil: ")
    Ganjil(o)

    fmt.Println()
}

func Ganjil(n int) {
    if n <= 0 {
        return
    }

    Ganjil(n - 1)

    if n%2 != 0 {
        fmt.Print(n, " ")
    }
}