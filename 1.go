package main
import "fmt"
func main() {
	var i, a, b, hasil int
	fmt.Scan(&a,&b)
	
	for i = 1; i <= b; i++ {
		hasil += a * i
	}	
	fmt.Print(hasil)
		
}