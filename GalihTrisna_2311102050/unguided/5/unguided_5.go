package main
import "fmt"

func ganjil(n int, i int) {
	if i > n {
		return
	} else {
		if i%2 != 0 {
			fmt.Print(i," ")
		}
		ganjil(n, i+1)
	}
}

func main() {
    var n int
    fmt.Scan(&n)
    ganjil(n,1)
}
