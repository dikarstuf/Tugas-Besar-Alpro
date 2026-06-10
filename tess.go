package main

import "fmt"

func main() {
	var char rune
	var hasil bool

	// Menggunakan fmt.Scanf dengan "%c" untuk membaca satu karakter
	fmt.Scanf("%c", &char)

	// Cek apakah karakter adalah huruf besar (A-Z)
	hurufbesar := (char >= 'A' && char <= 'Z')

	// Cek apakah karakter adalah huruf kecil (a-z)
	hurufkecil := (char >= 'a' && char <= 'z')

	// Hasil adalah true jika salah satu kondisi terpenuhi
	hasil = hurufbesar || hurufkecil

	fmt.Println(hasil)
}