package main
import "fmt"

func barisan(n int) {
    if n == 1 {
        fmt.Print("1")
        return
    }
    fmt.Print(n, " ")
    barisan(n - 1)  
    fmt.Print(" ", n) 
}

func main() {
    var n int
    fmt.Scan(&n)
    barisan(n)
}
