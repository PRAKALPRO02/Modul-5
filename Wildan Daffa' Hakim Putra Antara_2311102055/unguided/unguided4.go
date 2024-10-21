package main

import "fmt"

func printPattern(n, current int) {
	if current == 1 {
		fmt.Printf("%d ", current)
		return
	}
	fmt.Printf("%d ", current)
	printPattern(n, current-1)
	fmt.Printf("[%d+%d]", n, current-1)
	fmt.Printf("%d ", current)
}

func main() {
	var nVerse int
	fmt.Scan(&nVerse)
	printPattern(nVerse, nVerse)
}
