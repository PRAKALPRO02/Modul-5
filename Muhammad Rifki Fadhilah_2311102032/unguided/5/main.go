package main

import "fmt"

func urutanBilanganGanjil(n int) {
  if n < 1{
    return 
  }
  urutanBilanganGanjil(n - 1)
  if n % 2 == 1 {
    fmt.Print(n," ")
  }
}

func main() {
  var n int
  fmt.Scan(&n)
  urutanBilanganGanjil(n)
  fmt.Print()
}
