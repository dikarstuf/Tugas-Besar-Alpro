package main
import "fmt"
func main() {
	var n,hasil,i,x int
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&x)
		hasil = x * 10
		fmt.Println(hasil)
	}
	
		
}