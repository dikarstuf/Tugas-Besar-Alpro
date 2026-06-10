package main

import "fmt"

func main() {
	var buku, harga int
	var member string
	var total float64
	harga = 10000
	fmt.Scan(&buku, &member)
	total = float64(harga * buku)

	if member == "A" {
		if buku < 5 {
			total *= 0.90
		} else if buku >= 5 && buku <= 10 {
			total *= 0.80
		} else {
			total *= 0.70
		}

	}
	if member == "B" {
		if buku < 5 {
			total *= 0.95
		} else if buku >= 5 && buku <= 10 {
			total *= 0.90
		} else {
			total *= 0.85
		}

	}
	if member == "C" {
		if buku < 5 {
			total = total
		} else if buku >= 5 && buku <= 10 {
			total *= 0.95
		} else {
			total *= 0.90
		}

	}
	if member == "N" {
		if buku > 10 {
			total *= 0.95
		} else {
			total = total
		}

	}
	fmt.Print("RP ", total)
}
