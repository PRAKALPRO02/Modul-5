package main
import "fmt"


func turun(n int) {
	if n < 1 {
		return
	}

	fmt.Print(n, " ")
	turun(n - 1)

}


func naik(n, i int) {
	if i > n { // misal inputan 5 dan i saat ini 6 maka akan rekursif akan berhenti
		return
	}
	fmt.Print(i, " ")
	naik(n, i + 1)

	fmt.Println()

}


func main(){

	var n int

	fmt.Println("\n*PROGRAM MENGHITUNG BARIS BILANGAN*")
	fmt.Print("\nMasukkan angka : ")
	fmt.Scan(&n)

	fmt.Println()

	fmt.Println("Barisan bilangan : \n")
	turun(n)
	naik(n, 2)

	fmt.Println()

	
}