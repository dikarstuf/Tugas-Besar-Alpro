package main

import "fmt"

func main() {
	var digit int
	var d1, d2, d3 int
	var outp bool
	
	fmt.Scan(&digit)
	d3 = digit % 10
	d2 = (digit / 10) % 10
	d1 = digit / 100
	outp = (d1 > d2) && (d2 > d3)
	fmt.Println(outp)





}