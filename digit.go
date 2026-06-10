package main
import "fmt"
func main () {
	
	var x, d1, d2, d3 int
    
	fmt.Println("Masukkan nilai x: ")
	fmt.Scan(&x)
	d3 = x %10
	x = x / 10
	d2 = x %10
	x = x / 10
	d1 = x %10
	x = x / 10
	fmt.Println(d1, d2, d3)
}