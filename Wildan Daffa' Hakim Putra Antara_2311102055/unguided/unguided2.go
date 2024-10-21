package main

import "fmt"

func ColumnPrint(n, max int) {
	if n > max {
		return
	} else {
		rowPrint(1, n)
		fmt.Println()
		ColumnPrint(n+1, max)
	}

}

func rowPrint(current, max int) {
	if current > max {
		return
	} else {
		fmt.Print("*")
		rowPrint(current+1, max)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	ColumnPrint(1, n)
}
