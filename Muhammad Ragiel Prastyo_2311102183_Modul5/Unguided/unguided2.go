// MUHAMMAD RAGIEL PRASTYO
// 2311102183
package main
import (
    "fmt"
)

func cetakBintang(n int) {
    if n <= 0 {
        return
    }
    for i := 0; i < n; i++ {
        fmt.Print("*")
    }
    fmt.Println()
    cetakBintang(n - 1) 
}

func tampilkanPola(n int) {
    if n <= 0 {
        fmt.Println("Masukan harus lebih besar dari 0")
        return
    }
    cetakBintang(n)
}

func main() {
    var n int
    
    fmt.Print("Masukkan angka: ")
    fmt.Scan(&n) 
    
    fmt.Printf("\nPola bintang untuk N = %d:\n", n)
    tampilkanPola(n) 
}