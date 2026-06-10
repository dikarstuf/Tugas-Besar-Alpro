package main

import "fmt"

func main() {
	var D1,D2,D3,D4, A, B int
	var jumlah int
	fmt.Scan(&A)
	
	D1 = A % 10
	D2 = (A / 10) % 10
	D3 = (A / 100) % 10
	D4 = (A / 1000)
	
	B = (D4 * 1000) + (D1 * 100) + (D2 * 10) + D3 
	
	jumlah = A + B 
	
	fmt.Println(B, jumlah)
	
}