package main

import "fmt"

func cetakBintang(jumlah int){
	if jumlah <= 0 {
		return
	} 
	fmt.Print("*")
	cetakBintang(jumlah - 1)
}

func baris(n, i int) { // n untuk menentukan jumlah baris yang akan dicetak. 
	if i > n {
		return
	}
	cetakBintang(i) // akan mencetak bintang sebanyak i kali 
	fmt.Println()
	baris(n, i + 1)

}

/*
misal n = 5



*/

func main(){
	var n int

	fmt.Println("\n*PROGRAM MENAMPILKAN POLA BINTANG*")
	fmt.Print("\nMasukkan jumlah baris(n) : ")
	fmt.Scan(&n)
	fmt.Println("\nPola : \n")
	
	baris(n, 1)

	fmt.Println()
}