// MUHAMMAD RAGIEL PRASTYO
// 2311102183
package main
import (
    "fmt"
)

func cetakUrutan(n int, current int, turun bool) {
    fmt.Printf("%d ", current)

    if turun {
        if current > 1 {
            cetakUrutan(n, current-1, true)
        } else {
            cetakUrutan(n, current+1, false)
        }
    } else { 
        if current < n {
            cetakUrutan(n, current+1, false)
        }
    }
}

func main() {
    var n int
    
    fmt.Print("Masukkan bilangan positif: ")
    fmt.Scan(&n)
    
    if n <= 0 {
        fmt.Println("Mohon masukkan bilangan positif!")
        return
    }
    
    fmt.Printf("Barisan bilangan untuk N = %d adalah: ", n)
    cetakUrutan(n, n, true)
    fmt.Println() 
}