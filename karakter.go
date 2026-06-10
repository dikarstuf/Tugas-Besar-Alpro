package main

import "fmt"

func main() {
	var inputChar byte
	var kategori string

	fmt.Scanf("%c", &inputChar)

	if inputChar >= '0' && inputChar <= '9' {
		kategori = "Bilangan"
	} else if inputChar >= 'A' && inputChar <= 'Z' {
		kategori = "Huruf Besar"
	} else if inputChar >= 'a' && inputChar <= 'z' {
		kategori = "Huruf Kecil"
	} else {
		kategori = "Simbol"
	}

	fmt.Println(kategori)
}
