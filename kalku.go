package main

import "fmt"

func main() {
	var x1, x2, xkondisi int
	fmt.Scan(&x1, &x2, &xkondisi)
	switch xkondisi {
	case 1:
		fmt.Println(x1 + x2)
	case 2:
		fmt.Println(x1 - x2)
	case 3:
		fmt.Println(x1 * x2)

	case 4:
		if x2 == 0 {
			fmt.Println("tidak terdefinisi")
		} else {
			fmt.Println(float64(x1) / float64(x2))
		}

	case 5:
		if x2 == 0 {
			fmt.Println("tidak terdefinisi")
		} else {
			fmt.Println(x1 % x2)
		}

	case 6:
		if x2 == 0 {
			fmt.Println("tidak terdefinisi")
		} else {
			fmt.Println(x1 / x2)
		}
	default:
		fmt.Println("Operator Tidak Valid")
	}
}
