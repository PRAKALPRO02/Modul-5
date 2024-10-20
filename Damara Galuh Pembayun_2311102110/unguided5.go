package main

import "fmt"

// cetakBilanganGanjil adalah fungsi rekursif untuk mencetak bilangan ganjil
func cetakBilanganGanjil(n int) {
        // Kasus dasar: jika n kurang dari 1, tidak ada bilangan ganjil yang perlu dicetak
        if n < 1 {
                return
        }

        // Panggil fungsi cetakBilanganGanjil dengan n-2 untuk mencetak bilangan ganjil sebelumnya
        cetakBilanganGanjil(n - 2)

        // Cetak bilangan ganjil saat ini
        if n >= 1 {
                fmt.Print(n, " ")
        }
}

func main() {
        var n int

        fmt.Print("Masukkan bilangan: ")
        fmt.Scan(&n)

        fmt.Printf("Barisan bilangan ganjil dari 1 hingga %d: ", n)
        cetakBilanganGanjil(n)
        fmt.Println()
}