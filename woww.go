package main
import "fmt"

func main() {
	h := map[string]int{
		"senin": 0, "Senin": 0,
		"selasa": 1, "Selasa": 1,
		"rabu": 2, "Rabu": 2,
		"kamis": 3, "Kamis": 3,
		"jumat": 4, "Jumat": 4,
		"sabtu": 5, "Sabtu": 5,
		"minggu": 6, "Minggu": 6,
	}
	
	for {
		var d string
		var t, b, y int
		fmt.Scan(&d)
		if d == "Exit" {
			break
		}
		fmt.Scan(&t, &b, &y)
		
		i := h[d]
		w := 0
		
		for w < 2 {
			// Increment hari
			i = (i + 1) % 7
			
			// Hitung jumlah hari dalam bulan SEBELUM increment tanggal
			m := 31
			if b == 2 {
				if y%400 == 0 || (y%100 != 0 && y%4 == 0) {
					m = 29
				} else {
					m = 28
				}
			} else if b == 4 || b == 6 || b == 9 || b == 11 {
				m = 30
			}
			
			// Increment tanggal
			t++
			
			// Jika tanggal melebihi hari dalam bulan
			if t > m {
				t = 1
				b++
				if b > 12 {
					b = 1
					y++
				}
			}
			
			// Hitung hanya hari kerja (Senin-Jumat = 0-4)
			if i >= 0 && i <= 4 {
				w++
			}
		}
		
		fmt.Printf(">> passport bisa diambil pada tanggal %d bulan %d tahun %d\n", t, b, y)
	}
}