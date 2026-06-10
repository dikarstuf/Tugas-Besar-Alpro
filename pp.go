package main

import "fmt"

func main() {
	var p, l, luas int
	var hasil bool
	fmt.Scan(&p, &l)
	luas = (p * l) 
	hasil = luas % 2 == 1
	fmt.Println(luas, hasil)

}