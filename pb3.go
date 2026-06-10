package main

import "fmt"

func main() {
	var n int 
	
	fmt.Scan(&n) 
	
	var hasil int64 = 1

	for i := 1; i <= n; i++ {
		hasil = hasil * int64(i)
	}

	fmt.Println(hasil)
}