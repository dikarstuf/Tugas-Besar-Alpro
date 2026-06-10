package main

import "fmt"

func main() {
	var n, d, fD, c, t, y int
	var x bool

	fmt.Scan(&n, &y)
	c = 0
	t = n

	for t >= 10 {
		t = t / 10
	}
	fD = t

	for n > 0 {
		d = n % y
		fmt.Println(d)
		
		x = (d % 2) == 0
		
		if x && d > fD {
			c = c + 1
		}
		
		fmt.Println(x)
		n = n / y
	}

	fmt.Println(c)
}