package main

import "fmt"

func pangkat(numPangkat, pangkatke int) int {
	if pangkatke == 1 {
		return numPangkat
	}
	return numPangkat * pangkat(numPangkat, pangkatke-1)
}

func main() {
	var n, pngktke int
	fmt.Scan(&n, &pngktke)
	fmt.Println(pangkat(n, pngktke))
}
