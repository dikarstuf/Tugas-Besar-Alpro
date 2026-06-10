package main
import "fmt"
func f(x,y int) int {
var hasil int
hasil = 2 * x * y + 5
return hasil
}
func main() {
	var x,y,hasil int
	fmt.Scan(&x,&y)
	hasil = f(x,y)
	fmt.Print(hasil)
}
