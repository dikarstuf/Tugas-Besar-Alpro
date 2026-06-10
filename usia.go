package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	if x >= 0 && x <= 12 {
		fmt.Println("Anak-anak")

	} else if x >= 13 && x <= 17 {
		fmt.Println("Remaja")

	} else if x >= 18 && x <= 55 {
		fmt.Println("Dewasa")

	} else if x > 55 {
		fmt.Println("Lansia")
	}
}
