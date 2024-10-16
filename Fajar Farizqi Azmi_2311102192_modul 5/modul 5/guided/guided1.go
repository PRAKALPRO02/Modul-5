package main

import "fmt"

func main () {
	var n int 
	fmt.Scan(&n) // memnbaca input pengguna
	baris (n)    // memanggil fungsi rekursif 'baris'

}

func baris(bilangan int) {
	if bilangan== 1 {         // base case : jika bilangan sama dengan1
		fmt.Println(1)        // cetak angka 1
	}else{                    // jika bilkangan klebih besar dari 1
		fmt.Println(bilangan) // cetak bilangan saat ini
		baris (bilangan -1 )  // panggil fungsi baris '
	}
}
