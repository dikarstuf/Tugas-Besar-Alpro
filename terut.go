package main

import "fmt"

func main() {
	var digit, d1, d2, d3 int 
	var meningkat, menurun, hasil bool
	
	fmt.Scan(&digit)
	
	
	d3 = digit % 10
	d2 = (digit / 10) % 10
	d1 = digit / 100
	
	meningkat = (d1 > d2) && (d2 > d3)
	menurun = (d1 < d2) && (d2 < d3)
	
	hasil = meningkat || menurun
	
	fmt.Println(hasil)
}