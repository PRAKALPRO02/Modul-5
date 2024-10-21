package main
import "fmt"

func bilanganGanjil(n int) {
    if n <= 0 {
        return
    }
    bilanganGanjil(n - 1)
    if n%2 != 0 {
        fmt.Print(n," ")
    }
}

func main() {
    var N int
    fmt.Print("masukkan angka: ")
    fmt.Scan(&N)
    bilanganGanjil(N)
}