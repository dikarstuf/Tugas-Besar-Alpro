package main

import "fmt"

func main() {
	var poin, totalPoin, j int
	var stop bool

	fmt.Scan(&poin)
	totalPoin = 0
	stop = poin < 0 || poin > 3
	j = 0

	for !stop {
		j = j + 1
		totalPoin = totalPoin + poin
		fmt.Scan(&poin)
		stop = totalPoin >= 30 || poin < 0 || poin > 3
	}

	fmt.Printf("%d %d\n", totalPoin, j)
}