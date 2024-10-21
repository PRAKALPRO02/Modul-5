 package main

 import "fmt"

 func ganjil(n, i int){ // i batas maksimum bilangan yang akan diperiksa dan dicetak 
	if n > i { // jika n > i maka rekursif akan berhenti
		return
	}
	if n % 2 != 0{ // jika hasil bagi n dengan 2 adalah 0, maka bilangan tsb adalah ganjil, lalu tampilkan pada layar 
		fmt.Print(n, " ")
	}
	ganjil(n+1, i) // n + 1, n dimulai dari 0, 
	fmt.Println()
 }

 func main(){
	var i int

	fmt.Println("\n\n*PROGRAM MENCARI BILANGAN GANJIL*")
	fmt.Print("\nMasukkan jumlah bilangan: ")
	fmt.Scan(&i)

	fmt.Println("\nBilangan yang ganjil : \n")
	ganjil(0, i)
	
 }