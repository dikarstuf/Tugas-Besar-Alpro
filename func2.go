package main

import ("fmt")

func hitungBeratIdeal(umur int) float64 {
	beratKg := float64((umur * 2) + 8)
	return beratKg * 2.20462
}

func tentukanKategori(umur int, beratSaatIni float64) {
	beratIdealLb := hitungBeratIdeal(umur)
	batasBawah := 0.9 * beratIdealLb
	batasAtas := 1.1 * beratIdealLb

	var kategori string

	if beratSaatIni < batasBawah {
		kategori = "Kurang"
	} else if beratSaatIni >= batasBawah && beratSaatIni <= batasAtas {
		kategori = "Ideal"
	} else {
		kategori = "Obesitas"
	}

	fmt.Printf("%.2f lb\n", beratIdealLb)
	fmt.Printf("Kategori: %s\n", kategori)
}

func main() {
	var umur int
	var beratSaatIni float64

	fmt.Scan(&umur, &beratSaatIni)

	if umur >= 1 && umur <= 10 {
		tentukanKategori(umur, beratSaatIni)
	}
}