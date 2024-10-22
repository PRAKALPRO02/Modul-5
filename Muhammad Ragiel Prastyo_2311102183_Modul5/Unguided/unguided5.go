// MUHAMMAD RAGIEL PRASTYO
// 2311102183
package main
import (
    "fmt"
)

func cetakBilanganGanjil(current int, n int) {
    if current > n {
        return
    }

    fmt.Printf("%d ", current)

    cetakBilanganGanjil(current + 2, n)
}

func main() {
    var n int
    
    fmt.Print("Masukkan bilangan positif: ")
    fmt.Scan(&n)
    
    if n <= 0 {
        fmt.Println("Mohon masukkan bilangan positif!")
        return
    }
    
    fmt.Printf("Bilangan ganjil dari 1 hingga %d adalah: ", n)
    cetakBilanganGanjil(1, n)
    fmt.Println() 
}