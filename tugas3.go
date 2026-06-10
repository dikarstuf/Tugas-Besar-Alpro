package main
import "fmt"

func main() {
	// Deklarasi variabel
	var x int
	var y int
	var hasil float64

	x = 70
	y = 20

	hasil = 1.0/(3.0*float64(x)*float64(x)+10.0) + 10.0*float64(y)+7.0

	// Keluaran	
	fmt.Println(hasil)
}