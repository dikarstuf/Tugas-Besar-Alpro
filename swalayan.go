package main

import "fmt"

func main() {
	var a, y, c, d, e int
	var cas, dis float64
	var x string
	fmt.Scan(&x, &a, &y, &c, &d, &e)

	semuaGanjil := a%2 != 0 && y%2 != 0 && c%2 != 0 && d%2 != 0 && e%2 != 0
	semuaGenap := a%2 == 0 && y%2 == 0 && c%2 == 0 && d%2 == 0 && e%2 == 0
	cas1 := 3.1 * float64(a+y+c)
	dis1 := 1.7 * float64(c+d+e)

	if semuaGanjil {
		cas = 0.0
		dis = dis1
	} else if semuaGenap {
		cas = cas1
		dis = 0.0
	} else {
		cas = cas1
		dis = dis1
	}

	if x == "yes" {
		const memberBonus = 1.15
		cas *= memberBonus
		dis *= memberBonus
	}

	if cas > 35.00 {
		cas = 35.00
	}

	fmt.Printf("cash back %.2f, dan discount %.2f\n", cas, dis)
}
