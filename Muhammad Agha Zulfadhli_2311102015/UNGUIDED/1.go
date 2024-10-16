package main

import "fmt"

func rekursif(prev1, prev2, max, iterasi int) {
	if iterasi <= max {
		fmt.Print(prev2+prev1, " ")
		prev1 = prev2 + prev1
		rekursif(prev2, prev1, max, iterasi+1)
	}
}

func main() {
	var prev2, prev1 int = 0, 1
	var max int
	fmt.Print("Jumlah Iterasi : ")
	fmt.Scan(&max)
	rekursif(prev1, prev2, max, 1)
}
