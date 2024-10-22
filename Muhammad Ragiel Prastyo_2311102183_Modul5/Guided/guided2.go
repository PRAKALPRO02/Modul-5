// MUHAMMAD RAGIEL PRASTYO
// 2311102183
package main
import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(penjumlahan(n))
}

func penjumlahan(n int) int {
	if n == 1 {
		return 1
}	else {
	return n + penjumlahan(n -1)
}
}