package main

import "fmt"

func main() {
	var x, y int
	var status bool

	fmt.Scan(&x)
	y = x

	for status = false; !status; {
		fmt.Scan(&x)
		if x > y && x != 0 {
			y = x
		}
		status = x == 0
	}
	fmt.Println(y)
}
