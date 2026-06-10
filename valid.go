package main

import "fmt"

func validateNumber(x int) {
	var p3, p7, p5, p0, p2, p6 int = -1, -1, -1, -1, -1, -1
	var temp int = x
	var count int = 0
	var digit int
	var valid bool = true

	var cTemp int = x
	for cTemp > 0 {
		cTemp /= 10
		count++
	}

	temp = x
	for i := count; i > 0; i-- {
		digit = temp % 10
		if digit == 3 {
			p3 = i
		} else if digit == 7 {
			p7 = i
		} else if digit == 5 {
			p5 = i
		} else if digit == 0 {
			p0 = i
		} else if digit == 2 {
			p2 = i
		} else if digit == 6 {
			p6 = i
		}
		temp /= 10
	}

	if (p3 != -1 && p7 != -1 && p7 < p3) ||
		(p5 != -1 && p0 != -1 && p0 < p5) ||
		(p2 != -1 && p6 != -1 && p6 < p2) {
		valid = false
	}

	if valid {
		fmt.Println("BILANGAN VALID")
	} else {
		fmt.Println("BILANGAN TIDAK VALID")
	}
}

func main() {
	var input int
	fmt.Scan(&input)
	validateNumber(input)
}