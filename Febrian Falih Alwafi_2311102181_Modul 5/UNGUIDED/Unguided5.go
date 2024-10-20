package main

import "fmt"

func main() {
    fmt.Print("Masukan bilangan: ")
    var n int
    fmt.Scan(&n)
    cetakGanjil(n)
}

func cetakGanjil(n int) {
    if n <= 0 {
        return
    }//2311102181

    cetakGanjil(n - 1)

    if n%2 != 0 {
        fmt.Print(n, " ")
    }
}


