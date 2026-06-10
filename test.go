package main

import (
	"fmt"
)

func main() {
	var x, y float64
	nama := "Tim Dosen A1Pro_1"

	fmt.Print("Masukkan bilangan bulat x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan bulat y: ")
	fmt.Scan(&y)

	if 7+2*x == 0 || y == 0 {
		fmt.Printf("%s +Inf\n", nama)
		return
	}

	hasil := (22.0/y)/(7+2*x) + x*y

	fmt.Printf("%s %v\n", nama, hasil)
}
