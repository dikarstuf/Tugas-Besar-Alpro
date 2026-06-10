package main

import "fmt"

func main() {
	var c rune
	var count int

	fmt.Scanf("%c", &c)

	for c != '#' {
		count = count + 1

		fmt.Scanf("%c", &c)
	}

	fmt.Println(count)
}
