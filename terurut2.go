package main

import "fmt"

func main() {

	var digit, ds int
	var d1, d2, d3 int
	var outpf, outp1, outp2 bool

	ds := digit
	fmt.Scan(&ds)
	d3 = ds % 10
	d2 = ds % 10
	d1 = ds % 10

	outp1 = bool(d1 > d2) && (d2 > d3)
	outp2 = bool(d1 < d2) && (d2 < d3)
	outpf = outp1 || outp2
	fmt.Println(outpf)

}
