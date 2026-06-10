package main

import ("fmt")

func hitungJamBelajar(N float64, K float64, Wbelajar *float64) {
	var totalBelajar float64 = 0
	var waktuBermain float64

	for i := 1; i <= int(N); i++ {
		fmt.Scan(&waktuBermain)
		totalBelajar += (K - waktuBermain)
	}

	if N > 0 {
		*Wbelajar = totalBelajar / N
	} else {
		*Wbelajar = 0
	}
}

func main() {
	var n, k, wBelajar float64

	fmt.Scan(&n)
	fmt.Scan(&k)
	wBelajar = 0

	hitungJamBelajar(n, k, &wBelajar)

	fmt.Printf("%.1f\n", wBelajar)
}