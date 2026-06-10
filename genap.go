package main

import "fmt"

func main() {
	var b, tot, dig int
	tot = 0

	fmt.Scan(&b)

	for b > 0 {
		dig = b % 10

		tot = tot + (1 - (dig % 2))

		b = b / 10
	}

	fmt.Println(tot)
}
