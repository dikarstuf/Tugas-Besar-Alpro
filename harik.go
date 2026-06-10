package main

import (
	"fmt"
)

func hariKerja(hari int, sisa int) int {
	if sisa == 0 {
		return 0
	}

	kerja := 0
	if hari >= 1 && hari <= 5 {
		kerja = 1
	}

	nextHari := hari + 1
	if hari == 7 {
		nextHari = 1
	}

	return kerja + hariKerja(nextHari, sisa-1)
}

func main() {
	var hariMulai, jumlahHari int
	fmt.Scan(&hariMulai, &jumlahHari)

	fmt.Println(hariKerja(hariMulai, jumlahHari))
}