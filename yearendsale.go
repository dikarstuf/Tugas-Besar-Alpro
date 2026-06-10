package main

import (
	"fmt"
)

func main() {
	var transaksi int
	var potongan, finalTransaksi float64

	fmt.Print("Masukkan total transaksi: ")
	fmt.Scan(&transaksi)

	if transaksi > 1000000 {
		potongan = 0.2 * float64(transaksi)
	} else if transaksi >= 100000 {
		potongan = 0.15 * float64(transaksi)
	} else {
		potongan = 0.05 * float64(transaksi)
	}

	if potongan > 2000000 {
		potongan = 2000000
	}

	finalTransaksi = float64(transaksi) - potongan

	fmt.Printf("Potongan: %.0f\n", potongan)
	fmt.Printf("Final Transaksi: %.0f\n", finalTransaksi)
}