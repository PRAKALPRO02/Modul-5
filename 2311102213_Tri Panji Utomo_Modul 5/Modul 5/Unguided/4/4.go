package main

import "fmt"

func turun213(n int) {
    if n <= 0 {
        return
    }

    fmt.Print(n, " ")
    turun213(n - 1)
}

func naik213(n, awal int) {
    if n > awal {
        return
    }

    fmt.Print(n, " ")
    naik213(n+1, awal)
}

func main() {
    var o int

    fmt.Print("Masukkan bilangan bulat positif: ")
    fmt.Scan(&o)

    turun213(o)

    if o >= 1 {
        naik213(2, 0)
    }
	fmt.Println()
}