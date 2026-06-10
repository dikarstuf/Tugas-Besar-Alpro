package main

import "fmt"
func main() {
	var alice, bob,selisih int
	var hasil bool
	fmt.Scan(&alice, &bob)

	selisih = alice - bob
	
	if selisih < 0 {
		selisih = -selisih
	}
	
	hasil = alice == bob || selisih == 1 || selisih == 5
	
	fmt.Printf("Menang? %t\n", hasil)
}