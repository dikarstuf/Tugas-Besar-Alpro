package main

import "fmt"

func findRepunit(k int, length int, rem int) int {
	var nextRem int = (rem*10 + 1) % k

	if nextRem == 0 {
		return length
	}

	if length >= k {
		return -1
	}

	return findRepunit(k, length+1, nextRem)
}

func smallestRepunitDivByK(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}

	if 1%k == 0 {
		return 1
	}

	return findRepunit(k, 2, 1%k)
}

func main() {
	var k int
	var result int

	fmt.Scan(&k)

	result = smallestRepunitDivByK(k)
	fmt.Println(result)
}