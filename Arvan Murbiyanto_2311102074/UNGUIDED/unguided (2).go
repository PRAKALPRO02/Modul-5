package main
import "fmt"

func printbintangStr(n int) string {
	if n == 1 {
		return "*"
	} else {
		return printbintangStr(n-1) + "\n" + cetakBintang(n)
	}
}
func cetakBintang(n int) string {
	bintang := ""
	for i := 0; i < n; i++ {
		bintang += "*"
	}
	return bintang
}
func printbintangDirect(n int) {
	if n == 1 {
		fmt.Println("*")
		return
	}
	printbintangDirect(n - 1)
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}
func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)
	fmt.Println("\nHasil menggunakan string:")
	fmt.Println(printbintangStr(n))
	fmt.Println("\nHasil mencetak langsung:")
	printbintangDirect(n)
}
