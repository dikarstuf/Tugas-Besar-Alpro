package main
import "fmt"
func main() {
	var td,jam,menit,detik,sd int
	fmt.Scan(&td)
	jam = td / 3600
	sd = td % 3600 
	menit = sd  / 60
	detik = sd % 60
	fmt.Printf("%d jam, %d menit, %d detik" ,jam,menit,detik)
}