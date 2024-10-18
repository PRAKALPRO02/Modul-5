package main

import "fmt"

func ganjil(a int) {
	for u:=0; u <= a; u++ { 
	if u % 2 == 1 {
		fmt.Print(u," ")
	}
}
}
func main() {
	var a int
	fmt.Scan(&a)
	ganjil(a)
}