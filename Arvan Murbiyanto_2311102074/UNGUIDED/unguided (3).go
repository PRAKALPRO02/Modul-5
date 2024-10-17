package main
import "fmt"

func findFactors(num, divisor int) {
	if divisor > num {
		return
	}

	if num%divisor == 0 {
		fmt.Print(divisor, " ")
	}

	findFactors(num, divisor+1)
}
func cetakFaktor(n, i int) {
	if i > n {
		return
	}

	if n%i == 0 {
		fmt.Print(i, " ")
	}

	cetakFaktor(n, i+1)
}
func main() {
	var number int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&number)
	fmt.Printf("Faktor dari %d adalah (findFactors): ", number)
	findFactors(number, 1)
	fmt.Println()
	fmt.Printf("Faktor dari %d adalah (cetakFaktor): ", number)
	cetakFaktor(number, 1)
	fmt.Println()
}
