package main

import "fmt"

// pangkat adalah fungsi rekursif untuk menghitung x pangkat y
func pangkat(x, y int) int {
        // Kasus dasar: jika y adalah 0, maka hasil pangkatnya adalah 1
        if y == 0 {
                return 1
        }

        // Kasus rekursif: pangkat(x, y) = x * pangkat(x, y-1)
        return x * pangkat(x, y-1)
}

func main() {
        var x, y int

        fmt.Print("Masukkan bilangan dasar (x): ")
        fmt.Scan(&x)
        fmt.Print("Masukkan pangkat (y): ")
        fmt.Scan(&y)

        hasil := pangkat(x, y)
        fmt.Printf("%d pangkat %d adalah %d\n", x, y, hasil)
}