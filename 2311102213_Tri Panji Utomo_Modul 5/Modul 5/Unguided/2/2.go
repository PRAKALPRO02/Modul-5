package main

import "fmt"

func main() {
    var N int

    fmt.Print("Masukkan jumlah baris: ")
    fmt.Scan(&N)

    pola(N, 1)
}

func bintang213(n, i int) {
    if i == 0 {
        return
    }

    fmt.Print("*")
    bintang213(n, i-1)
}

func pola(n, baris int) {
    if baris > n {
        return
    }

    bintang213(n, baris)
    fmt.Println()
    pola(n, baris+1)
}