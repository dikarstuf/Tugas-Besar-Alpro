package main
import "fmt"
func main() {
	var utangA,utangB,jumlahlembarA,sisalembarA,jumlahlembarB,sisalembarB,uang int
	fmt.Scan(&utangA,&utangB)
	
	uang = 2000
	
	jumlahlembarA = utangA / uang
	sisalembarA = utangA % uang
	
	jumlahlembarB = utangB / uang
	sisalembarB = utangB % uang
	
	if sisalembarA > 0 {
		jumlahlembarA = jumlahlembarA + 1
	}
	if sisalembarB > 0 {
		jumlahlembarB = jumlahlembarB + 1
	}
	fmt.Printf("PT. Alice memerlukan %d lembar USD 2,000\n",jumlahlembarA)
	fmt.Printf("PT. Bob memerlukan %d lembar USD 2,000\n", jumlahlembarB)
}