// fajar farizqi azmi
// 2311102192

package main

import "fmt"

func printOddNumbers(n int) {
    if n < 1 {
        return // 2311102192
    }
    printOddNumbers(n - 2)
    if n % 2 != 0 {
        fmt.Print(n, " ")
    }
}

func main () {
    var number int
    fmt.Print("Masukkan bilangan: ")
    fmt.Scan(&number)

    fmt.Printf("Bilangan ganjil dari 1 sampai %d: ", number)
    printOddNumbers(number)
    fmt.Println()
}