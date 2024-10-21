package main
import "fmt"

func main()  {
	var n int
	fmt.Scan(&n)
	fmt.Println(penjumlah(n))
}

func penjumlah(n int) int  {
	if n == 1 {
		return 1
	} else {
		return n + penjumlah(n-1)
	}
}

