package main

import "fmt"

// bintang adalah fungsi rekursif untuk mencetak pola bintang
func bintang(n int) {
        // Kasus dasar: jika n = 1, cetak satu bintang
        if n == 1 {
                fmt.Println("*")
                return
        }

        // Panggil fungsi bintang dengan n-1 untuk mencetak baris sebelumnya
        bintang(n - 1)

        // Cetak baris saat ini dengan n bintang
        for i := 0; i < n; i++ {
                fmt.Print("*")
        }
        fmt.Println()
}

func main() {
        var n int

        fmt.Print("Masukkan jumlah baris: ")
        fmt.Scan(&n)

        // Panggil fungsi bintang untuk mencetak pola
        bintang(n)
}