// fajar farizqi azmi
// 2311102192


package main

import "fmt"

func printSequence(n int) {
    if n == 0 {
        return // 2311102192
    }
    fmt.Print(n, " ")
    printSequence(n - 1)
    if n != 1 {
        fmt.Print(n, " ")
    }
}

func main() {
    var number int
    fmt.Print("Masukkan bilangan: ")
    fmt.Scan(&number)

    fmt.Printf("Barisan bilangan: ")
    printSequence(number)
    fmt.Println()
}