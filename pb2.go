package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
    
	hasil := fmt.Sprint(n)

	for i := n - 1; i >= 1; i-- {
		hasil = hasil + " x " + fmt.Sprint(i)
	}

	fmt.Println(hasil)
}