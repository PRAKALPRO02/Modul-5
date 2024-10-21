package main
import "fmt"

func fibonacci (Sn int) int {
	if Sn<=1{
		return Sn
	}
	return fibonacci(Sn-1) + fibonacci(Sn-2)
}

func main (){
	var jumlahDeret int
	fmt.Print("masukkan nilai hingga suku ke n berapa, yang ingin ditampilkan: ")
	fmt.Scan(&jumlahDeret)

	fmt.Print("\nSn:")
	for i:=0;i<jumlahDeret+1;i++{
		fmt.Printf(" %d ",fibonacci(i))
	}

}