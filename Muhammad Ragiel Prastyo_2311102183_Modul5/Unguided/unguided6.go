// MUHAMMAD RAGIEL PRASTYO
// 2311102183
package main
import (
    "fmt"
)

func pangkat(basis int, eksponen int) int {
    if eksponen == 0 {
        return 1
    }

    if eksponen == 1 {
        return basis
    }

    return basis * pangkat(basis, eksponen-1)
}

func main() {
    var x, y int
    
    fmt.Print("Masukkan dua bilangan (x y): ")
    fmt.Scan(&x, &y)
    
    if y < 0 {
        fmt.Println("Pangkat tidak boleh negatif!")
        return
    }
    
    hasil := pangkat(x, y)
    fmt.Printf("%d dipangkatkan %d = %d\n", x, y, hasil)
}