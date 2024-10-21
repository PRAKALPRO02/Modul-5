package main
import "fmt"

func Pangkat(x, y int) int {
	if y == 0 {
		return 1
	}
	return x * Pangkat(x, y-1)
}
   
func main() {
	var x int
	var y int

	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)

	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	hasile := Pangkat(x, y)
	fmt.Printf("%d dipangkatkan %d adalah %d\n", x, y, hasile)
}
