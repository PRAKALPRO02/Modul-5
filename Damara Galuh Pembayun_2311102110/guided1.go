package main

import "fmt"

func main (){
	var int 
	fmt.Scan (&n) //membaca input pengguna
	baris (n) //memanggil fungsi reskrusif 'baris'
}
func baris (bilangan int){
	if bilangan == 1 { //basecase apabila bilangan sama dengan 1
		fmt.Println (1) //cetak angka 1
	} else { //jika bilangan lebih besar dari 1
		fm.Prinln (bilangan) //cetak bilangan saat ini
		baris (nilangan-1) //panggil fungsi baris dengan bilangan 1
	}
}