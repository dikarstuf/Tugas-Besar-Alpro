package main
import "fmt"
func main() {
	var n,i,hasil,bil int
	fmt.Scan(&n)
	for i = 1; i <= n; i++ { 
	fmt.Scan(&bil)
	hasil = bil * i
	fmt.Print(hasil," ")
	}


}