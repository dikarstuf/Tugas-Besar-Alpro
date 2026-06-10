package main

import "fmt"

func main() {
	var angka1, angka, i int
	var equal bool

	fmt.Scan(&angka1)

	fmt.Scan(&angka)
	for angka != 0 {
		i++
		if angka == angka1 {
			equal = true
			break
		}
		fmt.Scan(&angka)

	}

	if equal {
		fmt.Print(i)
	} else {
		fmt.Println("Tidak Ada")
	}

}
