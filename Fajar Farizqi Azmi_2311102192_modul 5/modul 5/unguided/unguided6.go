// fajar farizqi azmi
// 2311102192

package main

import "fmt"

func power(x, y int) int {
    if y == 0 {
        return 1 // 2311102192
    }
    return x * power(x, y-1)
}

func main () {
    var base, exponent int
    fmt.Print("Masukkan bilangan pokok (x) : ")
    fmt.Scan(&base)
    fmt.Print("Masukkan pangkat (y) : ")
    fmt.Scan(&exponent)

    result := power(base, exponent)
    fmt.Printf("%d pangkat %d adalah : %d\n", base, exponent, result)
}