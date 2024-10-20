package main 

import  "fmt"


func fibo (n int) int {
	if n == 0{
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibo (n-1) + fibo (n-2)
	}
}

func main () {
	var bilangan int

	fmt.Print ("\nDeret Fibonacci hingga suku ke-n = ")
	fmt.Scan (&bilangan)

	for n :=0; n <= bilangan; n++ {
		fmt.Printf ("%d ", fibo(n) )
	}

	fmt.Println()

}