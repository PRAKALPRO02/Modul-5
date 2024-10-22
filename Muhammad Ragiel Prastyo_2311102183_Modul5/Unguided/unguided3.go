// MUHAMMAD RAGIEL PRASTYO
// 2311102183
package main
import (
    "fmt"
)

func findFactors(n int, current int) {
    if current > n {
        return
    }

    if n%current == 0 {
        fmt.Printf("%d ", current)
    }

    findFactors(n, current+1)
}

func main() {
    var n int
    
    fmt.Print("Masukkan bilangan positif: ")
    fmt.Scan(&n)
    
    if n <= 0 {
        fmt.Println("Mohon masukkan bilangan positif!")
        return
    }
    
    fmt.Printf("Faktor dari %d adalah: ", n)
    findFactors(n, 1)
    fmt.Println() 
}