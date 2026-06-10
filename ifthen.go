package main
import "fmt"
func main() {
	var a,b,c int
	fmt.Scan(&a,&b,&c)
	if a<0 {
		a = -a
	}
	if b<0 {
		b = -b
	}
	if c<0 {
		c = -c
	}
	fmt.Print(a,b,c)
}