package main
import "fmt"
func reverseNumber(num int) int {
	var hasil int = 0
	for num > 0 {
		hasil = hasil*10 + num%10
		num = num / 10
	}
	return hasil
}

func ambilGabung3Digit(num int) int {
	var tigaDigit, balik3, depan int
	balik3 = 0
	tigaDigit = num % 1000

	for tigaDigit > 0 {
		balik3 = balik3*10 + tigaDigit%10
		tigaDigit = tigaDigit / 10
	}

	depan = num / 1000
	return depan*1000 + balik3
}

func prosesData(n int) {
	var rev int
	var hasilAkhir int
	rev = reverseNumber(n)
	hasilAkhir = ambilGabung3Digit(rev)
	fmt.Println(hasilAkhir)
	cekDigit(hasilAkhir)
}

func cekDigit(x int) {
	var jumlahGanjil, jumlahGenap, digit int
	jumlahGanjil = 0
	jumlahGenap = 0

	for x > 0 {
		digit = x % 10
		if digit%2 == 0 {
			jumlahGenap++
		} else {
			jumlahGanjil++
		}
		x = x / 10
	}

	if jumlahGanjil > jumlahGenap {
		fmt.Println("Valid")
	} else {
		fmt.Println("Invalid")
	}
}

func main() {
	var input int
	fmt.Scan(&input)
	prosesData(input)
}