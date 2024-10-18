package main

import "fmt"

func faktorBilangan(n,i int){
	 if i > n {
    return
   }
   if n % i == 0{
    fmt.Print(i, " ")
   }
   faktorBilangan(n, i+1)
}

func main() {
	var n int
	fmt.Scan(&n)
  faktorBilangan(n,1)
  fmt.Print()
}