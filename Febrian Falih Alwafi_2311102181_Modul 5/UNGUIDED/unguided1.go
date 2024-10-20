package main

import "fmt"

func fibonacci(n int) int {
    if n <= 1 {
        return n
    }

    a, b := 0, 1
    for i := 2; i <= n; i++ {
        a, b = b, a+b
    }
    return b
}

func main() {
    var n int
    fmt.Print("Masukkan nilai n: ")
    fmt.Scanln(&n)

    for i := 0; i <= n; i++ {
        fmt.Printf("Fibonacci(%d) = %d\n", i, fibonacci(i))
    } //2311102181
}



