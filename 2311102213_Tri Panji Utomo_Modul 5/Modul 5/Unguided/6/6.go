package main

import "fmt"

func pnkt213(x, y int) int {
    if y == 0 {
        return 1
    }
    return x * pnkt213(x, y-1)
}

func main() {
    var x, y int

    fmt.Print("Masukkan bilangan x: ")
    fmt.Scan(&x)

    fmt.Print("Masukkan bilangan y: ")
    fmt.Scan(&y)

    hasil := pnkt213(x, y)
    fmt.Printf("%d dipangkatkan %d adalah %d\n", x, y, hasil)
}