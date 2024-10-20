package main

import "fmt"

// faktor adalah fungsi rekursif untuk mencari faktor dari suatu bilangan
func faktor(bilangan, pembagi int) {
        // Kasus dasar: jika pembagi lebih besar dari bilangan, tidak ada faktor lagi
        if pembagi > bilangan {
                return
        }

        // Jika bilangan habis dibagi pembagi, maka pembagi adalah faktor
        if bilangan%pembagi == 0 {
                fmt.Print(pembagi, " ")
        }

        // Panggil fungsi faktor dengan pembagi berikutnya
        faktor(bilangan, pembagi+1)
}

func main() {
        var bilangan int

        fmt.Print("Masukkan bilangan: ")
        fmt.Scan(&bilangan)

        fmt.Printf("Faktor dari %d adalah: ", bilangan)
        faktor(bilangan, 1) // Memulai pencarian faktor dari 1
        fmt.Println()
}