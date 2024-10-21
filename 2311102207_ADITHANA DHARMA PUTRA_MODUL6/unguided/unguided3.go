package main
import "fmt"

func faktor(n, i int) {
    if i > n {
        return
    }
    if n%i == 0 {
        fmt.Print(i," ")
    }
    faktor(n, i+1)
}

func main() {
    var N int
    fmt.Print("masukan angka: ")
    fmt.Scan(&N)
    fmt.Print("faktor dari ",N, "= ")
    faktor(N, 1)
}