package main

import "fmt"

func main() {
	var x, n, s, stot, i float64
	
	fmt.Scan(&x, &n)
	
	for i = 1; i <= n; i++ {
		s = (n - i + 1) / (i * x) 
		stot += s
	}
	
	fmt.Printf("%.3f\n", stot)
}