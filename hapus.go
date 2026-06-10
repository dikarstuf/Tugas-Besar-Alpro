package main

import "fmt"

func jumlahDigit(n int) int {
	if n == 0 {
		return 0
	}
	return 1 + jumlahDigit(n/10)
}

func hapusGanjil(n, curr, del int) int {
	if n == 0 {
		return 0
	}
	
	digit := n % 10
	sisa := n / 10
	
	if curr == del {
		return hapusGanjil(sisa, curr+1, del)
	}
	
	return hapusGanjil(sisa, curr+1, del)*10 + digit
}

func hapusGenap(n, curr, del1, del2 int) int {
	if n == 0 {
		return 0
	}
	
	digit := n % 10
	sisa := n / 10
	
	if curr == del1 || curr == del2 {
		return hapusGenap(sisa, curr+1, del1, del2)
	}
	
	return hapusGenap(sisa, curr+1, del1, del2)*10 + digit
}

func main() {
	var n int
	fmt.Scan(&n)

	panjang := jumlahDigit(n)

	if panjang <= 2 {
		fmt.Println("INPUT ERROR")
	} else if panjang%2 != 0 {
		// Posisi dihapus dari belakang (1-based)
		// Rumus: panjang - (panjang + 1)/2 + 1
		posisiDel := panjang - (panjang+1)/2 + 1
		fmt.Println(hapusGanjil(n, 1, posisiDel))
	} else {
		// Posisi dihapus dari belakang (1-based)
		p := panjang / 2
		posisiDel2 := panjang - p + 1      // P+1 dari depan
		posisiDel1 := panjang - (p + 1) + 1 // P dari depan
		fmt.Println(hapusGenap(n, 1, posisiDel1, posisiDel2))
	}
}