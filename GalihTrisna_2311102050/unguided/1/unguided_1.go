package main
import "fmt"

func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func deretFobonacci(n int, i int) {
    if i > n {
        return
    } else {
        fmt.Print(" ",fibonacci(i))
        deretFobonacci(n, i+1)
    }
}

func main() {
    var n int
    fmt.Scan(&n)
    deretFobonacci(n, 0)
}
