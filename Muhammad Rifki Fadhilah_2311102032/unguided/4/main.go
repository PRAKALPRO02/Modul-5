package main

import "fmt"

func main() {
  var n int
  fmt.Scan(&n)
  urutanBilangan(n)
  fmt.Print()
}

func urutanBilangan(n int) {
	if n < 1{
		return 
	}
	fmt.Print(n," ")
	urutanBilangan(n - 1)
  if n > 1 {
    fmt.Print(n," ")
  }
}