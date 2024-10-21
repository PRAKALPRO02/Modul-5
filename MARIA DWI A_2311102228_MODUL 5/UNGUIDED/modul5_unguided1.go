package main

import "fmt"

/*
func fibonaci(n int) {
	var f1, f2, f3, int

	f2 = 0
	f3 = 1
	for i := 1; i <= n; i++{
		fmt.Println(3)
		f1 = f2
		f2 = f3
		f3 = f1 + f2
	}
}
	*/ 

func fibonacci(n int) int {
	if (n == 0 || n == 1){
		return n
	} else {
		return (fibonacci(n-1) + (fibonacci(n-2)))
	}
}

func main(){
	var n int
	j:= 0

	fmt.Println("\n*PROGRAM MENGHITUNG DERET FIBONACCI*")
	fmt.Print("\nBatas deret: ")
	fmt.Scan(&n)
	fmt.Println()

	for i:= 0; i <= n; i++{
		fmt.Print("n", i, " : ")
		fmt.Println(fibonacci(j))
		j++
	}
	fmt.Println()
}