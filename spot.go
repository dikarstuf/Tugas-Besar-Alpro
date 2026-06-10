package main
import (
	"fmt"
)

const NMAX = 999

type song struct {
	judul, penyanyi         string
	durasiMenit, durasiDetik int
}

type TabLagu struct {
	totalDurasi int
	arrLagu     [NMAX]song
}

func main() {
	var playlist TabLagu
	var n int

	fmt.Scan(&n)

	if n > 0 {
		inputLagu(&playlist, n)
		printLagu(playlist, n)
	}
}

func inputLagu(s *TabLagu, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Print("Judul lagu: ")
		fmt.Scan(&s.arrLagu[i].judul)
		
		fmt.Print("Nama Penyanyi: ")
		fmt.Scan(&s.arrLagu[i].penyanyi)
		
		fmt.Print("Durasi lagu (menit detik): ")
		fmt.Scan(&s.arrLagu[i].durasiMenit, &s.arrLagu[i].durasiDetik)
		
		s.totalDurasi += (s.arrLagu[i].durasiMenit * 60) + s.arrLagu[i].durasiDetik
		fmt.Println()
	}
}

func printLagu(s TabLagu, n int) {
	var line string = "+----------------------+--------------+----------+"
	var header string = "| Judul Lagu           | Penyanyi     | Durasi   |"
	var i int
	var jam, sisa, menit, detik int
	
	fmt.Println(line)
	fmt.Println(header)
	fmt.Println(line)

	for i = 0; i < n; i++ {
		fmt.Printf("| %-20s | %-12s | %02d:%02d    |\n", 
			s.arrLagu[i].judul, 
			s.arrLagu[i].penyanyi, 
			s.arrLagu[i].durasiMenit, 
			s.arrLagu[i].durasiDetik)
	}

	fmt.Println(line)

	jam = s.totalDurasi / 3600
	sisa = s.totalDurasi % 3600
	menit = sisa / 60
	detik = sisa % 60

	fmt.Printf("| Total Durasi Lagu                   | %02d:%02d:%02d |\n", 
		jam, menit, detik)
	fmt.Println(line)
}