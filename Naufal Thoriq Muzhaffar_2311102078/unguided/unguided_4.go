package main
import "fmt"

func cetakBarisan(n, current int) {
    if current == 1 {
        fmt.Print(current, " ")
        return
    }
    fmt.Print(current, " ")
    cetakBarisan(n, current-1)
    fmt.Print(current, " ")
}

func main() {
    var n int
    fmt.Print("Masukkan bilangan bulat positif N: ")
    fmt.Scan(&n)
    cetakBarisan(n, n)
}
