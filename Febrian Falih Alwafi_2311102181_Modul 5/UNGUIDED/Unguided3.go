package main

import "fmt"

func main() {
    var n int
    fmt.Print("Masukkan bilangan bulat positif: ")
    fmt.Scanln(&n)

    fmt.Println("Faktor dari", n, "adalah:", faktor(n))
}

func faktor(n int) []int {
    var faktor []int
    for i := 1; i <= n; i++ {
        if n%i == 0 {
            faktor = append(faktor, i)
        }
    }
    return faktor
}



