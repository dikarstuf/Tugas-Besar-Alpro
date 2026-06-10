package main

import "fmt"

func main() {
	var x, i int
	var prima bool

	fmt.Scan(&x)

	prima = true

	if x <= 1 {
		prima = false
	} else {
		for i = 2; i*i <= x; i++ {
			if x%i == 0 {
				prima = false
				break
			}
		}
	}

	fmt.Println(prima)
}
