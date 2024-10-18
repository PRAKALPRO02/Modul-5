package main

import "fmt"


func bintang(baris, barisSekarang, b int) {
    if barisSekarang > baris {
        return
    }
    if b < barisSekarang {
        fmt.Print("*")
        bintang(baris, barisSekarang, b+1)
    } else {
        fmt.Println()
        bintang(baris, barisSekarang+1, 0)
    }
}

func main() {
    var baris int
	fmt.Scan(&baris)
    bintang(baris, 1, 0)
}
