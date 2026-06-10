package main

import "fmt"

func main() {
	var A,B,C,d,dd,kuadrat int
	fmt.Scan(&d)

	A = d / 100
	B = (d % 100) / 10
	C = d % 10

	dd += A * 100000 + A * 10000
	dd += B * 1000 + B * 100
	dd += C * 10 + C * 1

	kuadrat = dd * dd

	fmt.Println(dd)
	fmt.Println(kuadrat)
}