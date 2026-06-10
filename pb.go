package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var x, y int

		fmt.Scan(&x, &y)
		fmt.Println(x * y)
	}
}