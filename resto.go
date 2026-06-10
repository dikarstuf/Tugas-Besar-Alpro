package main

import ("fmt")

func hitungPoin(belanja int) int {
	if belanja < 50000 {
		return 0
	} else if belanja < 100000 {
		return 5
	} else if belanja < 200000 {
		return 10
	}
	return 20
}

func prosesTransaksi(no int, totalPoin int) int {
	var belanja int
	fmt.Scan(&belanja)

	if belanja == 0 {
		return totalPoin
	}

	poin := hitungPoin(belanja)

	fmt.Printf("Transaksi %d: Belanja Rp %d, Poin: %d\n", no, belanja, poin)

	return prosesTransaksi(no+1, totalPoin+poin)
}

func main() {
	total := prosesTransaksi(1, 0)
	fmt.Println("Total Poin:", total)
}