package main

import ("fmt")

func countDigit(n int) int {
	if n == 0 {
		return 0
	}
	return 1 + countDigit(n/10)
}

func getNumber(n int, i int) int {
	if i == 1 {
		return n % 10
	}
	return getNumber(n/10, i-1)
}

func checkFirstLast(n int) bool {
	jumlahDigit := countDigit(n)
	first := getNumber(n, jumlahDigit)
	last := getNumber(n, 1)
	return first == last
}

func sumDigit(n int) int {
	if n == 0 {
		return 0
	}
	return (n % 10) + sumDigit(n/10)
}

func isPrime(n int, i int) bool {
	if n <= 1 {
		return false
	}
	if i*i > n {
		return true
	}
	if n%i == 0 {
		return false
	}
	return isPrime(n, i+1)
}

func main() {
	var N int
	fmt.Scan(&N)

	jumlahDigit := countDigit(N)

	tengahIndex := (jumlahDigit + 1) / 2
	digitTengah := getNumber(N, jumlahDigit-tengahIndex+1)

	jumlah := sumDigit(N)

	if checkFirstLast(N) && isPrime(digitTengah, 2) && jumlah%3 == 0 {
		fmt.Println("Pesawat merupakan pesawat spesial.")
	} else if !checkFirstLast(N) {
		fmt.Println("Pesawat merupakan pesawat berisiko.")
	} else if !isPrime(digitTengah, 2) {
		fmt.Println("Pesawat merupakan pesawat biasa.")
	} else if jumlah%3 != 0 {
		fmt.Println("Pesawat tidak layak terbang.")
	}
}