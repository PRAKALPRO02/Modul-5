package main
import "fmt"

func main() {
    var n int
    fmt.Print("Masukkan nilai n: ")
    fmt.Scan(&n)
    polaBintang(n, 1)
}

func polaBintang(n, current int) {
    if current > n {
        return
    }
    cetakBintang(current)
    fmt.Println()
    polaBintang(n, current+1)
}

func cetakBintang(jumlah int) {
    if jumlah == 0 {
        return
    }
    fmt.Print("*")
    cetakBintang(jumlah - 1)
}
