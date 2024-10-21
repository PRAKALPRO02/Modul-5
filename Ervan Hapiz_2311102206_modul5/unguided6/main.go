package main

import "fmt"

func pangkat(n, i int) int {
	if i == 1 {
		return n
	} else {
		return n * pangkat(n, i-1)
	}
}
func main() {
	var n, i int
	fmt.Scan(&n, &i)
	fmt.Print("Hasil pangkat : ", pangkat(n, i))

}
