package main

import "fmt"

// cetakBarisan adalah fungsi rekursif untuk mencetak barisan bilangan
func cetakBarisan(n int) {
        // Kasus dasar: jika n = 1, cetak 1
        if n == 1 {
                fmt.Print(n, " ")
                return
        }

        // Cetak bilangan saat ini
        fmt.Print(n, " ")

        // Panggil fungsi cetakBarisan dengan n-1 untuk mencetak bilangan sebelumnya
        cetakBarisan(n - 1)

        // Cetak bilangan saat ini lagi (saat kembali dari rekursi)
        fmt.Print(n, " ")
}

func main() {
        var n int

        fmt.Print("Masukkan bilangan: ")
        fmt.Scan(&n)

        fmt.Printf("Barisan bilangan dari %d hingga 1 dan kembali: ", n)
        cetakBarisan(n)
        fmt.Println()
}