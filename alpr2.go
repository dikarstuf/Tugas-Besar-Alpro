package main

import "fmt"

func main() {
	var n int
	var biner string

	fmt.Scan(&n)

	if n == 0 {
		biner = "0"
	}

	for n > 0 {
		if n%2 == 0 {
			biner = "0" + biner
		} else {
			biner = "1" + biner
		}
		n = n / 2
	}

	fmt.Println(biner)
}
