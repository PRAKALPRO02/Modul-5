package main

import "fmt"

func pangkat (x int, y int) int {
	if y == 0 {
		return 1
	}
	if y < 0 {
		return 1 / pangkat(x, -y)
	}
	return x * pangkat(x, y-1)
}

func main(){
	var x, y int
	fmt.Print("MAsukan bilangan bulat: ")
	fmt.Scan(&x)
	fmt.Print("Masukan Pangkat: ")
	fmt.Scan(&y)
	hasil := pangkat(x, y)
	fmt.Print("Hasil ", x," Pangkat ", y," Adalah: ", hasil )
}