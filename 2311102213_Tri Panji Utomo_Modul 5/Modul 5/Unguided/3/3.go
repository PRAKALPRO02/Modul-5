package main

import "fmt"

func main() {
    var bil213 int

    fmt.Print("Masukkan bilangan: ")
    fmt.Scan(&bil213)

    fmt.Print("Faktor dari ", bil213, " adalah: ")

    Faktor(bil213, bil213)
    fmt.Println()
}

func Faktor(bil213, pembagi int) {
    if pembagi == 1 {
        fmt.Print(1)
        return
    }

    if pembagi == 1 {
        return
    }

    if bil213%pembagi == 0 {
        fmt.Print(pembagi, " ")
    }

    Faktor(bil213, pembagi-1)
}