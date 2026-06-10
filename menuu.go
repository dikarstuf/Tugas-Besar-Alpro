package main

import "fmt"

type Menu struct {
	Nama   string
	Harga  int
	Stok   int
	Status string
}

const MAX int = 999

type arrMenu [MAX]Menu

func inputDataMenu(m *arrMenu, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Print("Nama Menu: ")
		fmt.Scan(&m[i].Nama)
		fmt.Print("Harga Menu: Rp.")
		fmt.Scan(&m[i].Harga)
		fmt.Print("Jumlah Stok: ")
		fmt.Scan(&m[i].Stok)

		if m[i].Stok > 0 {
			m[i].Status = "Tersedia"
		} else {
			m[i].Status = "Kosong"
		}
	}
}

func printDataMenu(m arrMenu, n int) {
	var i int
	fmt.Printf("\n%-15s %-12s %-12s %-15s\n", "Nama Menu", "Harga (Rp)", "Stok (pcs)", "Status Kesediaan")
	for i = 0; i < n; i++ {
		fmt.Printf("%-15s Rp.%-9d %-12d %-15s\n", m[i].Nama, m[i].Harga, m[i].Stok, m[i].Status)
	}
}

func main() {
	var n int
	var daftarMenu arrMenu

	fmt.Scan(&n)

	if n > MAX {
		n = MAX
	}

	inputDataMenu(&daftarMenu, n)
	printDataMenu(daftarMenu, n)
}