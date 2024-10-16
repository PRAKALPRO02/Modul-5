// fajar farizqi azmi
// 2311102192

package main

import "fmt"

func printbintang(n int) {
    if n == 1 {
        fmt.Println("*")
        return // 2311102192
    }
    printbintang(n - 1)
    for i := 0; i < n; i++ {
        fmt.Print("*")
    }
    fmt.Println()
}

func main() {
    var n int
    fmt.Print("Masukkan nilai N: ")
    fmt.Scan(&n)

    printbintang(n)
}