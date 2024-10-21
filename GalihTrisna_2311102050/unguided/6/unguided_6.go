package main
import "fmt"

func pangkat(x int, y int) int {
	if y == 1 {
		return x
	} else {
		return x * pangkat(x,y-1)
	}
}

func main() {
  	var x, y int
	fmt.Scanln(&x, &y)
	println(pangkat(x,y))
}
