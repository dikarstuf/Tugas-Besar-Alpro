package main

import "fmt"

const NMAX int = 50

type Jadwal struct {
	id         string
	lapangan   string
	jamMulai   int
	jamSelesai int
	status     string
	penyewa    string
}

type Riwayat struct {
	idJadwal string
	lapangan string
	jamMain  string
	penyewa  string
}

type tJadwal [NMAX]Jadwal
type tRiwayat [NMAX]Riwayat

var dataRiwayat tRiwayat
var nRiwayat int

func initData(J *tJadwal, n *int) {
	*n = 6
	J[0] = Jadwal{"01", "Lapangan A", 14, 15, "Kosong", "-"}
	J[1] = Jadwal{"02", "Lapangan B", 19, 20, "Kosong", "-"}
	J[2] = Jadwal{"03", "Lapangan A", 9, 10, "Kosong", "-"}
	J[3] = Jadwal{"04", "Lapangan A", 21, 22, "Dipesan", "kesmoy FC"}
	J[4] = Jadwal{"05", "Lapangan B", 16, 17, "Kosong", "-"}
	J[5] = Jadwal{"06", "Lapangan B", 10, 11, "Dipesan", "yessir"}
}

// Selection Sort Ascending
func cetakTabel(J tJadwal, n int) {

	var temp tJadwal = J

	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if temp[j].jamMulai < temp[minIdx].jamMulai {
				minIdx = j
			}
		}

		temp[i], temp[minIdx] = temp[minIdx], temp[i]
	}

	fmt.Println("\n-------------------------------------------------------------------------")
	fmt.Printf("| %-6s | %-12s | %-11s | %-9s | %-15s |\n", "ID", "Lapangan", "Jam Main", "Status", "Penyewa")
	fmt.Println("-------------------------------------------------------------------------")

	for i := 0; i < n; i++ {
		fmt.Printf("| %-6s | %-12s | %02d:00-%02d:00 | %-9s | %-15s |\n",
			temp[i].id, temp[i].lapangan, temp[i].jamMulai, temp[i].jamSelesai, temp[i].status, temp[i].penyewa)
	}
	fmt.Println("-------------------------------------------------------------------------")
}

// Insertion Sort Descending
func urutJadwal(J *tJadwal, n int, R *tRiwayat, nR *int) {
	var temp tJadwal
	var m int = 0

	for i := 0; i < n; i++ {
		if J[i].status == "Kosong" {
			temp[m] = J[i]
			m++
		}
	}

	if m == 0 {
		fmt.Println("\nTidak ada jadwal kosong saat ini.")
		return
	}

	for i := 1; i < m; i++ {
		key := temp[i]
		j := i - 1

		for j >= 0 && temp[j].jamMulai < key.jamMulai {
			temp[j+1] = temp[j]
			j--
		}
		temp[j+1] = key
	}

	fmt.Println("\n>>> DAFTAR JADWAL KOSONG (INSERTION SORT DESCENDING) <<<")

	fmt.Println("-------------------------------------------------------------------------")
	fmt.Printf("| %-6s | %-12s | %-11s | %-9s | %-15s |\n", "ID", "Lapangan", "Jam Main", "Status", "Penyewa")
	fmt.Println("-------------------------------------------------------------------------")
	for i := 0; i < m; i++ {
		fmt.Printf("| %-6s | %-12s | %02d:00-%02d:00 | %-9s | %-15s |\n",
			temp[i].id, temp[i].lapangan, temp[i].jamMulai, temp[i].jamSelesai, temp[i].status, temp[i].penyewa)
	}
	fmt.Println("-------------------------------------------------------------------------")
}

// Binary Search
func cariJadwalIdx(J tJadwal, n int, idCari string) int {
	low := 0
	high := n - 1

	for low <= high {
		mid := (low + high) / 2
		if J[mid].id == idCari {
			return mid
		} else if J[mid].id < idCari {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func bookingLapangan(J *tJadwal, n int) {
	var idTarget, namaTim string

	cetakTabel(*J, n)

	fmt.Print("Pilih ID Jadwal yang mau dipesan: ")
	fmt.Scan(&idTarget)

	idx := cariJadwalIdx(*J, n, idTarget)

	if idx != -1 {
		if J[idx].status == "Kosong" {
			fmt.Print("Masukkan Nama Tim Anda: ")
			fmt.Scan(&namaTim)

			J[idx].status = "Dipesan"
			J[idx].penyewa = namaTim

			jamFormat := fmt.Sprintf("%02d:00-%02d:00", J[idx].jamMulai, J[idx].jamSelesai)

			dataRiwayat[nRiwayat] = Riwayat{
				idJadwal: J[idx].id,
				lapangan: J[idx].lapangan,
				jamMain:  jamFormat,
				penyewa:  namaTim,
			}
			nRiwayat = nRiwayat + 1

			fmt.Println("Booking sukses!")
		} else {
			fmt.Println("Lapangan sudah penuh")
		}
	} else {
		fmt.Println("ID Jadwal tidak ditemukan")
	}
}

func cetakRiwayat(J tJadwal, n int) {
	if nRiwayat == 0 {
		fmt.Println("\nBelum ada riwayat pesanan baru")
		return
	}

	var kataCari string
	fmt.Print("Masukkan Nama Penyewa/Tim yang dicari: ")
	fmt.Scan(&kataCari)

	ketemu := false
	headerDicetak := false

	//Sequential Search
	for i := 0; i < nRiwayat; i++ {
		if dataRiwayat[i].penyewa == kataCari {
			if !headerDicetak {
				fmt.Println("\n-------------------------------------------------------------------------")
				fmt.Printf("| %-12s | %-15s | %-15s | %-15s |\n", "ID Jadwal", "Lapangan", "Jam Main", "Penyewa")
				fmt.Println("-------------------------------------------------------------------------")
				headerDicetak = true
			}
			fmt.Printf("| %-12s | %-15s | %-15s | %-15s |\n",
				dataRiwayat[i].idJadwal, dataRiwayat[i].lapangan, dataRiwayat[i].jamMain, dataRiwayat[i].penyewa)
			ketemu = true
		}
	}

	if ketemu {
		fmt.Println("-------------------------------------------------------------------------")
	} else {
		fmt.Printf("\nRiwayat pesanan untuk penyewa '%s' tidak ditemukan.\n", kataCari)
	}
}

func main() {
	var dataJadwal tJadwal
	var nJadwal int
	var menuUtama int

	nRiwayat = 0
	initData(&dataJadwal, &nJadwal)

	for {
		fmt.Println("\n=====================================")
		fmt.Println("            MENU UTAMA               ")
		fmt.Println("=====================================")
		fmt.Println("1. Lihat semua jadwal")
		fmt.Println("2. Jadwal kosong")
		fmt.Println("3. Booking lapangan")
		fmt.Println("4. Cari Riwayat Pesanan")
		fmt.Println("5. Keluar")
		fmt.Println("=====================================")

		fmt.Print("Pilih menu : ")
		fmt.Scan(&menuUtama)

		if menuUtama == 1 {
			cetakTabel(dataJadwal, nJadwal)
		} else if menuUtama == 2 {
			urutJadwal(&dataJadwal, nJadwal, &dataRiwayat, &nRiwayat)
		} else if menuUtama == 3 {
			bookingLapangan(&dataJadwal, nJadwal)
		} else if menuUtama == 4 {
			cetakRiwayat(dataJadwal, nJadwal)
		} else if menuUtama == 5 {
			fmt.Println("Program selesai")
			break
		} else {
			fmt.Println("Menu tidak tersedia")
		}
	}
}
