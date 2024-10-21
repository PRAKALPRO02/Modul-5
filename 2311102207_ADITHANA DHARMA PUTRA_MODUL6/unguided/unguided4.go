package main
import "fmt"

func cetakBaris(nilai_N,lembah int) {
    if nilai_N > lembah {
    fmt.Print(nilai_N," ")
    cetakBaris(nilai_N-1,lembah)
    }
	fmt.Print(nilai_N," ")
    if nilai_N < lembah {
	return
    }

}

func main() {
    var nilai_N int
    fmt.Print("Masukkan nilai N: ")
    fmt.Scan(&nilai_N)
    cetakBaris(nilai_N, 1)
}