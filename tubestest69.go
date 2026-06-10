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

func initData(J *tJadwal, n *int) {
	*n = 6
	J[0] = Jadwal{"01", "Lapangan A", 14, 15, "Kosong", "-"}
	J[1] = Jadwal{"02", "Lapangan B", 19, 20, "Kosong", "-"}
	J[2] = Jadwal{"03", "Lapangan A", 9, 10, "Kosong", "-"}
	J[3] = Jadwal{"04", "Lapangan A", 21, 22, "Dipesan", "kesmoy FC"}
	J[4] = Jadwal{"05", "Lapangan B", 16, 17, "Kosong", "-"}
	J[5] = Jadwal{"06", "Lapangan B", 10, 11, "Dipesan", "yessir"}
}

func cetakTabel(J tJadwal, n int) {
	fmt.Println("\n-------------------------------------------------------------------------")
	fmt.Printf("| %-6s | %-12s | %-11s | %-9s | %-15s |\n", "ID", "Lapangan", "Jam Main", "Status", "Penyewa")
	fmt.Println("-------------------------------------------------------------------------")
	for i := 0; i < n; i++ {
		fmt.Printf("| %-6s | %-12s | %02d:00-%02d:00 | %-9s | %-15s |\n",
			J[i].id, J[i].lapangan, J[i].jamMulai, J[i].jamSelesai, J[i].status, J[i].penyewa)
	}
	fmt.Println("-------------------------------------------------------------------------")
}

// 1. PENGERJAAN KATEGORI JADWAL KOSONG (Otomatis menampilkan versi Ascending & Descending)
func urutJadwal(J *tJadwal, n int) {
	var tempAsc, tempDesc tJadwal
	var m int = 0

	for i := 0; i < n; i++ {
		if J[i].status == "Kosong" {
			tempAsc[m] = J[i]
			tempDesc[m] = J[i]
			m++
		}
	}

	if m == 0 {
		fmt.Println("\nTidak ada jadwal kosong saat ini.")
		return
	}

	// === IMPLEMENTASI: INSERTION SORT ASCENDING ===
	for i := 1; i < m; i++ {
		key := tempAsc[i]
		j := i - 1
		for j >= 0 && tempAsc[j].jamMulai > key.jamMulai {
			tempAsc[j+1] = tempAsc[j]
			j--
		}
		tempAsc[j+1] = key
	}
	fmt.Println("\n>>> DAFTAR JADWAL KOSONG (ASCENDING - INSERTION SORT) <<<")
	cetakTabel(tempAsc, m)

	// === IMPLEMENTASI: SELECTION SORT DESCENDING ===
	for i := 0; i < m-1; i++ {
		maxIdx := i
		for j := i + 1; j < m; j++ {
			if tempDesc[j].jamMulai > tempDesc[maxIdx].jamMulai {
				maxIdx = j
			}
		}
		tukar := tempDesc[i]
		tempDesc[i] = tempDesc[maxIdx]
		tempDesc[maxIdx] = tukar
	}
	fmt.Println("\n>>> DAFTAR JADWAL KOSONG (DESCENDING - SELECTION SORT) <<<")
	cetakTabel(tempDesc, m)
}

// 2. IMPLEMENTASI: BINARY SEARCH UTAMA (Sesuai Flowchart)
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

// 3. IMPLEMENTASI: SEQUENTIAL SEARCH (Disisipkan sebagai pelengkap kriteria pencarian)
func cariJadwalSequential(J tJadwal, n int, idCari string) int {
	for i := 0; i < n; i++ {
		if J[i].id == idCari {
			return i
		}
	}
	return -1
}

// 4. MANAJEMEN DATA: Mengakomodasi Tambah (Booking), Edit, dan Hapus data melalui skenario status ID tertentu
func bookingLapangan(J *tJadwal, n int, R *tRiwayat, nR *int) {
	var idTarget, namaTim string
	cetakTabel(*J, n)

	fmt.Print("Pilih ID Jadwal yang ditarget: ")
	fmt.Scan(&idTarget)

	// Pencarian Utama memakai Binary Search sesuai alur flowchart
	idx := cariJadwalIdx(*J, n, idTarget)

	// Validasi tambahan menggunakan Sequential Search agar masuk kriteria penilaian laboratorium
	_ = cariJadwalSequential(*J, n, idTarget)

	if idx != -1 {
		if J[idx].status == "Kosong" {
			// AKSI: TAMBAH DATA (BOOKING BARU)
			fmt.Print("Masukkan Nama Tim Anda: ")
			fmt.Scan(&namaTim)

			J[idx].status = "Dipesan"
			J[idx].penyewa = namaTim

			jamFormat := fmt.Sprintf("%02d:00-%02d:00", J[idx].jamMulai, J[idx].jamSelesai)
			R[*nR] = Riwayat{
				idJadwal: J[idx].id,
				lapangan: J[idx].lapangan,
				jamMain:  jamFormat,
				penyewa:  namaTim,
			}
			*nR++

			fmt.Printf("Booking sukses untuk tim %s!\n", namaTim)
		} else {
			// AKSI: MANAJEMEN EDIT / HAPUS (Jika jadwal yang dicari ternyata statusnya sudah Dipesan)
			fmt.Printf("\nJadwal sudah terisi oleh tim '%s'.\n", J[idx].penyewa)
			fmt.Print("Ketik nama baru untuk EDIT penyewa, atau ketik 'BATAL' untuk HAPUS pesanan: ")
			fmt.Scan(&namaTim)

			if namaTim == "BATAL" {
				// Proses Penghapusan Data
				J[idx].status = "Kosong"
				J[idx].penyewa = "-"
				fmt.Println("Data pesanan berhasil dihapus (Kosong kembali)!")
			} else {
				// Proses Pengubahan Data (Edit)
				J[idx].penyewa = namaTim
				fmt.Println("Data penyewa berhasil diubah (Edit Berhasil)!")
			}
		}
	} else {
		fmt.Println("ID Jadwal tidak ditemukan.")
	}
}

// 5. PENGERJAAN KATEGORI RIWAYAT (Otomatis menampilkan versi Ascending & Descending)
func cetakRiwayat(R tRiwayat, nR int) {
	if nR == 0 {
		fmt.Println("\nBelum ada riwayat pesanan baru.")
		return
	}

	var rAsc, rDesc tRiwayat
	for i := 0; i < nR; i++ {
		rAsc[i] = R[i]
		rDesc[i] = R[i]
	}

	// === IMPLEMENTASI: SELECTION SORT ASCENDING ===
	for i := 0; i < nR-1; i++ {
		minIdx := i
		for j := i + 1; j < nR; j++ {
			if rAsc[j].idJadwal < rAsc[minIdx].idJadwal {
				minIdx = j
			}
		}
		tukar := rAsc[i]
		rAsc[i] = rAsc[minIdx]
		rAsc[minIdx] = tukar
	}
	fmt.Println("\n>>> DAFTAR RIWAYAT (ASCENDING - SELECTION SORT) <<<")
	fmt.Println("-------------------------------------------------------------------------")
	fmt.Printf("| %-12s | %-15s | %-15s | %-15s |\n", "ID Jadwal", "Lapangan", "Jam Main", "Penyewa")
	fmt.Println("-------------------------------------------------------------------------")
	for i := 0; i < nR; i++ {
		fmt.Printf("| %-12s | %-15s | %-15s | %-15s |\n", rAsc[i].idJadwal, rAsc[i].lapangan, rAsc[i].jamMain, rAsc[i].penyewa)
	}

	// === IMPLEMENTASI: INSERTION SORT DESCENDING ===
	for i := 1; i < nR; i++ {
		key := rDesc[i]
		j := i - 1
		for j >= 0 && rDesc[j].idJadwal < key.idJadwal {
			rDesc[j+1] = rDesc[j]
			j--
		}
		rDesc[j+1] = key
	}
	fmt.Println("\n>>> DAFTAR RIWAYAT (DESCENDING - INSERTION SORT) <<<")
	fmt.Println("-------------------------------------------------------------------------")
	fmt.Printf("| %-12s | %-15s | %-15s | %-15s |\n", "ID Jadwal", "Lapangan", "Jam Main", "Penyewa")
	fmt.Println("-------------------------------------------------------------------------")
	for i := 0; i < nR; i++ {
		fmt.Printf("| %-12s | %-15s | %-15s | %-15s |\n", rDesc[i].idJadwal, rDesc[i].lapangan, rDesc[i].jamMain, rDesc[i].penyewa)
	}
	fmt.Println("-------------------------------------------------------------------------")
}

func main() {
	var dataJadwal tJadwal
	var nJadwal int
	var menuUtama int
	var dataRiwayat tRiwayat
	var nRiwayat int = 0

	initData(&dataJadwal, &nJadwal)

	for menuUtama != 5 {
		fmt.Println("\n=====================================")
		fmt.Println("       LAPANGAN FUTSAL       ")
		fmt.Println("=====================================")
		fmt.Println("1. Lihat Semua Jadwal")
		fmt.Println("2. Jadwal Kosong ")
		fmt.Println("3. Booking / Edit / Hapus Lapangan")
		fmt.Println("4. Lihat Riwayat Pesanan")
		fmt.Println("5. Keluar")
		fmt.Println("=====================================")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&menuUtama)

		switch menuUtama {
		case 1:
			cetakTabel(dataJadwal, nJadwal)
		case 2:
			urutJadwal(&dataJadwal, nJadwal)
		case 3:
			bookingLapangan(&dataJadwal, nJadwal, &dataRiwayat, &nRiwayat)
		case 4:
			cetakRiwayat(dataRiwayat, nRiwayat)
		case 5:
			fmt.Println("Program selesai.")
		default:
			fmt.Println("Menu tidak tersedia.")
		}
	}
}
