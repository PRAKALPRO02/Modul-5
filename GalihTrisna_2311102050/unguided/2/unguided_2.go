package main
import "fmt"

func bintang(n int) {
    if n == 0 {
        return
    }
    fmt.Print("*")
    bintang(n - 1)
}

func baris(n int, i int) {
    if i > n {
        return
    }
    bintang(i)
    fmt.Println()
    baris(n, i+1)
}

func main() {
    var n int
    fmt.Scan(&n) 
    baris(n, 1)
}
