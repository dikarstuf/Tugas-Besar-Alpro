package main

import "fmt"

func validSolution(x int) bool {
	var count [10]int
	var digit int
	var i int

	for i = 0; i < 9; i++ {
		digit = x % 10
		if digit == 0 {
			return false
		}
		count[digit]++
		x = x / 10
	}

	if x != 0 {
		return false
	}

	for i = 1; i <= 9; i++ {
		if count[i] != 1 {
			return false
		}
	}

	return true
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(validSolution(n))
	
}