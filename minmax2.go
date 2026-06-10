package main

import "fmt"

type Item struct {
	nama  string
	harga int
}

func main() {
	var N int
	fmt.Scan(&N)

	var items [100]Item
	var namaTermahal, namaTermurah string
	var maxHarga, minHarga int

	for i := 0; i < N; i++ {
		fmt.Scan(&items[i].nama)
		fmt.Scan(&items[i].harga)

		if i == 0 {
			maxHarga = items[i].harga
			minHarga = items[i].harga
			namaTermahal = items[i].nama
			namaTermurah = items[i].nama
		} else {
			if items[i].harga > maxHarga {
				maxHarga = items[i].harga
				namaTermahal = items[i].nama
			}
			if items[i].harga < minHarga {
				minHarga = items[i].harga
				namaTermurah = items[i].nama
			}
		}
	}

	fmt.Printf("%-16s %s\n", "Nama Item", "Harga")
	for i := 0; i < N; i++ {
		fmt.Printf("%-16s %d\n", items[i].nama, items[i].harga)
	}

	fmt.Println()
	fmt.Printf("Item termahal %s\n", namaTermahal)
	fmt.Printf("Item termurah %s\n", namaTermurah)
}