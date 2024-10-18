package main

import "fmt"

func main() {
  var x,y int
  fmt.Scan(&x,&y)
  fmt.Print(bilanganPangkat(x,y))
}

func bilanganPangkat(x,y int) int {
  if y  == 0{
    return 1
  }else{
    return x * bilanganPangkat(x,y-1)
  }

}

