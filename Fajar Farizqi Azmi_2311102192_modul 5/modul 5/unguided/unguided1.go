// fajar farizqi azmi
// 2311102192

package main

import "fmt"

func fibonacci(n int) int {
    if n <= 1 {
        return n // 2311102192
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
    var n int
    fmt.Print("Masukkan nilai n: ") 
    fmt.Scan(&n)

    for i := 0; i <= n; i++ {
        fmt.Printf("Fibonacci ke-%d: %d\n", i, fibonacci(i))
    }
}