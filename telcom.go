package main

import ("fmt")

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func hitungDigitRek(nomor string, posisi int) int {
	if posisi >= len(nomor) {
		return 0
	}
	if isDigit(nomor[posisi]) {
		return 1 + hitungDigitRek(nomor, posisi+1)
	}
	return hitungDigitRek(nomor, posisi+1)
}

func awalanValid(nomor string) bool {
	if len(nomor) >= 2 && nomor[0:2] == "08" {
		return true
	}
	if len(nomor) >= 3 && nomor[0:3] == "+62" {
		return true
	}
	return false
}

func validasiNomor(nomor string) bool {
	if !awalanValid(nomor) {
		return false
	}

	start := 0
	if len(nomor) >= 2 && nomor[0:2] == "08" {
		start = 2
	} else if len(nomor) >= 3 && nomor[0:3] == "+62" {
		start = 3
	}

	jumlahDigit := hitungDigitRek(nomor, start)

	if jumlahDigit >= 9 && jumlahDigit <= 12 {
		return true
	}
	return false
}

func main() {
	var nomor string
	fmt.Scan(&nomor)

	if validasiNomor(nomor) {
		fmt.Println("VALID")
	} else {
		fmt.Println("TIDAK VALID")
	}
}